#include <stdio.h>

// eva
#include "adt/sds.h"
#include "rng/srng64.h"
#include "rng/srng64_normal.h"

// mlvm
#include "vm.h"

// cmd
#include "../helpers.h"
#include "mnist.h"

static error_t initModelWeight(struct vm_t*, struct srng64_t*,
                               struct opopt_t* rng, int w);
static error_t prepareData(struct srng64_t* seed, float32_t* x_data,
                           size_t x_size, float32_t* y_data, size_t y_size);

#define IMAGE_SIZE (28 * 28)
#define LABEL_SIZE (10)

#define FAKE_DATA 1

static unsigned char* images = NULL;
static unsigned char* labels = NULL;

int
main()
{
        error_t          err  = OK;
        sds_t            s    = sdsEmpty();
        struct srng64_t* seed = srng64New(123);

        {
                // x[bs, is]    -- is = IMAGE_SIZE
                // y[bs, ls]    -- ls = LABEL_SIZE
                //
                // forward pass
                //
                //   z[1] = zeros([1])
                //
                //   h1[bs, h1]   = matmul(x[bs, is], w1[is, h1])
                //   h1b[bs, h1]  = h1[bs, h1] + b1[h1]
                //   z1[bs, h1]   = max(h1b[bs, h1], z[1])
                //
                //   h2[bs, h2]   = matmul(z1[bs, h1], w2[h1, h2])
                //   h2b[bs, h2]  = h2[bs, h2] + b2[h2]
                //   z2[bs, h2]   = max(h2b[bs, h2], z[1])
                //
                //   o[bs, ls]    = matmul(z2[bs, h2], w3[h2, ls])
                //   l[bs]        = softmax_cross_entropy_with_logits(
                //                      y[bs, ls], o[bs, ls])
                //   loss[1]      = sum(l[bs])
                //
                //  backward pass
                //
                //   d_o[bs, ls]  = grad_softmax_cross_entropy_with_logits(
                //                     y[bs, ls], o[bs, ls])
                //   d_w3[h2, ls] = matmul(z2[bs, h2], d_o[bs, ls], trans_a)
                //   d_z2[bs, h2] = matmul(o[bs, ls], w3[h2, ls], trans_b)
                //
                //   -- the second matmul
                //   state_0      = cmpL(h2b[bs, h2], z[1])
                //   d_h2b[bs, h2]= mul(d_z2[bs, h2], state_0)
                //
                //   d_h2[bs, h2] = d_h2b[bs, h2]
                //   d_b2[h2]     = sum(d_h2b[bs, h2], axis=1)
                //
                //   d_w2[h1, h2] = matmul(z1[bs, h1], d_h2[bs, h2], trans_a)
                //   d_z1[bs, h1] = matmul(d_h2[bs, h2], w2[h1, h2] trans_b)
                //
                //   -- the first matmul
                //   state_1      = cmpL(h1b[bs, h1], z[1])
                //   d_h1b[bs, h1]= mul(d_z1[bs, h1], state1)
                //   d_h1[bs, h1] = d_h1b[bs, h1]
                //   d_b1[h1]     = sum(d_h1b[bs, h1], axis=1)
                //   d_w1[is, h1])= matmul(x[bs, is], d_h1[bs, h1], trans_a)
        }

        const int bs   = 32;
        const int is   = IMAGE_SIZE;
        const int ls   = LABEL_SIZE;
        const int h1_s = 64;
        const int h2_s = 64;

        // ---
        // defines vm with some shapes.
        struct vm_t*   vm = vmNew();
        struct opopt_t opt;

        struct shape_t* sp_x      = R2S(vm, bs, is);
        struct shape_t* sp_y      = R2S(vm, bs, ls);
        struct shape_t* sp_w1     = R2S(vm, is, h1_s);
        struct shape_t* sp_h1     = R2S(vm, bs, h1_s);
        struct shape_t* sp_b1     = R1S(vm, h1_s);
        struct shape_t* sp_w2     = R2S(vm, h1_s, h2_s);
        struct shape_t* sp_h2     = R2S(vm, bs, h2_s);
        struct shape_t* sp_b2     = R1S(vm, h2_s);
        struct shape_t* sp_w3     = R2S(vm, h2_s, ls);
        struct shape_t* sp_o      = R2S(vm, bs, ls);
        struct shape_t* sp_l      = R1S(vm, bs);
        struct shape_t* sp_scalar = R1S(vm, 1);

        int x    = vmTensorNew(vm, F32, sp_x);
        int y    = vmTensorNew(vm, F32, sp_y);
        int z    = vmTensorNew(vm, F32, sp_scalar);
        int w1   = vmTensorNew(vm, F32, sp_w1);
        int h1   = vmTensorNew(vm, F32, sp_h1);
        int b1   = vmTensorNew(vm, F32, sp_b1);
        int h1b  = vmTensorNew(vm, F32, sp_h1);
        int z1   = vmTensorNew(vm, F32, sp_h1);
        int w2   = vmTensorNew(vm, F32, sp_w2);
        int b2   = vmTensorNew(vm, F32, sp_b2);
        int h2   = vmTensorNew(vm, F32, sp_h2);
        int h2b  = vmTensorNew(vm, F32, sp_h2);
        int z2   = vmTensorNew(vm, F32, sp_h2);
        int w3   = vmTensorNew(vm, F32, sp_w3);
        int o    = vmTensorNew(vm, F32, sp_o);
        int l    = vmTensorNew(vm, F32, sp_l);
        int loss = vmTensorNew(vm, F32, sp_scalar);

        // ---
        // init weights
        {
                printf("init model weights.\n");
                opt.mode = 0;  // std normal.
                NE(initModelWeight(vm, seed, &opt, w1));
                NE(initModelWeight(vm, seed, &opt, b1));
                NE(initModelWeight(vm, seed, &opt, w2));
                NE(initModelWeight(vm, seed, &opt, b2));
                NE(initModelWeight(vm, seed, &opt, w3));
        }

        // ---
        // fetch inputs.
        {
                float32_t *x_data, *y_data;
                NE(vmTensorData(vm, x, (void**)&x_data));
                NE(vmTensorData(vm, y, (void**)&y_data));
                NE(prepareData(seed, x_data, /*x_size=*/sp_x->size, y_data,
                               /*y_size=*/sp_y->size));
        }

        // ---
        // forward pass
        {
                struct oparg_t prog[] = {
                    {OP_MATMUL, h1, x, w1, 0},
                };
                vmBatch(vm, 1, prog);
                NE(vmExec(vm, OP_MATMUL, NULL, h1, x, w1));
                NE(vmExec(vm, OP_ADD, NULL, h1b, h1, b1));
                NE(vmExec(vm, OP_MAX, NULL, z1, h1b, z));
                NE(vmExec(vm, OP_MATMUL, NULL, h2, z1, w2));
                NE(vmExec(vm, OP_ADD, NULL, h2b, h2, b2));
                NE(vmExec(vm, OP_MAX, NULL, z2, h2b, z));
                NE(vmExec(vm, OP_MATMUL, NULL, o, z2, w3));
                NE(vmExec(vm, OP_LS_SCEL, NULL, l, y, o));
                OPT_SET_REDUCTION_SUM(opt, 0);
                NE(vmExec(vm, OP_REDUCE, &opt, loss, l, -1));

                S_PRINTF("logits: ", o, "\n");
                S_PRINTF("labels: ", y, "\n");
                S_PRINTF("loss after softmax cel: ", l, "\n");
                S_PRINTF("loss: ", loss, "\n");
                printf("%s\n", s);
        }

cleanup:
        if (images != NULL) free(images);
        if (labels != NULL) free(labels);
        srng64Free(seed);
        vmFree(vm);
        sdsFree(s);
        return err;
}

// impl
static error_t
readMnistData(unsigned char** images, unsigned char** labels)
{
        error_t err = readMnistTrainingImages(images);
        if (err) {
                return err;
        }

        err = readMnistTrainingLabels(labels);
        if (err) {
                return err;
        }
        printf("sample label %d -- image:\n", (int)**labels);
        printMnistImage(*images);
        printf("smaple label %d -- image:\n", (int)*(*labels + 1));
        printMnistImage(*images + 28 * 28);
        return OK;
}

static void
prepareFakeData(struct srng64_t* seed, float32_t* x_data, size_t x_size,
                float32_t* y_data, size_t y_size)
{
        srng64StdNormalF(seed, x_size, x_data);
        srng64StdNormalF(seed, y_size, y_data);
}

error_t
prepareData(struct srng64_t* seed, float32_t* x_data, size_t x_size,
            float32_t* y_data, size_t y_size)
{
        if (FAKE_DATA) {
                printf("generating fake minis data.");
                prepareFakeData(seed, x_data, x_size, y_data, y_size);
                return OK;
        } else {
                error_t err;
                printf("reading real minis data.");
                if ((err = readMnistData(&images, &labels))) {
                        if (images != NULL) free(images);
                        if (labels != NULL) free(labels);
                        return err;
                }
                return OK;
        }
}

error_t
initModelWeight(struct vm_t* vm, struct srng64_t* seed, struct opopt_t* opt,
                int w)
{
        struct srng64_t* weight_seed = srng64Split(seed);
        opt->rng_seed                = *weight_seed;
        error_t err                  = vmExec(vm, OP_RNG, opt, w, -1, -1);
        srng64Free(weight_seed);

        return err;
}
