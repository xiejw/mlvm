#include <stdio.h>

#include "adt/sds.h"
#include "rng/srng64.h"
#include "rng/srng64_normal.h"

#include "vm.h"

#include "mnist.c"

static error_t readMnistData(unsigned char** images, unsigned char** labels);
static void    prepareFakeData(struct srng64_t* seed, float32_t* x_data,
                               size_t x_size, float32_t* y_data, size_t y_size);
static error_t initModelWeight(struct vm_t*, struct srng64_t*,
                               struct opopt_t* rng, int w);

#define IMAGE_SIZE (28 * 28)
#define LABEL_SIZE (10)

#define FAKE_DATA 1

#define NO_ERR(e) NO_ERR_IMPL_(e, __FILE__, __LINE__)

#define NO_ERR_IMPL_(e, f, l)                                           \
        if (e) {                                                        \
                err = e;                                                \
                errDump("failed to exec op @ file %s line %d\n", f, l); \
                goto cleanup;                                           \
        }

#define R1S(vm, s1)     vmShapeNew(vm, 1, (int[]){(s1)});
#define R2S(vm, s1, s2) vmShapeNew(vm, 2, (int[]){(s1), (s2)});

int main()
{
        error_t          err    = OK;
        unsigned char*   images = NULL;
        unsigned char*   labels = NULL;
        struct srng64_t* seed   = srng64New(123);

        // x[bs, is]    -- is = IMAGE_SIZE
        // y[bs, ls]    -- ls = LABEL_SIZE
        //
        // forward pass
        //
        //   z[1] = zeros([1])
        //
        //   h1[bs, h1]  = matmul(x[bs, is], w1[is, h1])
        //   h1b[bs, h1] = h1[bs, h1] + b1[h1]
        //   z1[bs, h1]  = max(h1b[bs, h1], z[1])
        //
        //   h2[bs, h2]  = matmul(z1[bs, h1], w2[h1, h2])
        //   h2b[bs, h2] = h2[bs, h2] + b2[h2]
        //   z2[bs, h2]  = max(h2[bs, h2], z[1])
        //
        //   o [bs, ls]  = matmul(z2[bs, h2], w3[h2, ls])
        //   loss[1]     = softmax_cross_entropy_with_logits(
        //                     y[bs, ls], o[bs, ls])

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
        int loss = vmTensorNew(vm, F32, sp_scalar);

        // ---
        // init weights
        opt.mode = 0;  // std normal.

        printf("init model weights.\n");
        NO_ERR(initModelWeight(vm, seed, &opt, w1));
        NO_ERR(initModelWeight(vm, seed, &opt, b1));
        NO_ERR(initModelWeight(vm, seed, &opt, w2));
        NO_ERR(initModelWeight(vm, seed, &opt, b2));
        NO_ERR(initModelWeight(vm, seed, &opt, w3));

        // ---
        // fetch inputs.
        float32_t *x_data, *y_data;
        {
                NO_ERR(vmTensorData(vm, x, (void**)&x_data));
                NO_ERR(vmTensorData(vm, y, (void**)&y_data));

                if (FAKE_DATA) {
                        printf("generating fake minis data.");
                        prepareFakeData(seed, x_data, /*x_size=*/sp_x->size,
                                        y_data,
                                        /*y_size=*/sp_y->size);
                } else {
                        printf("reading real minis data.");
                        if ((err = readMnistData(&images, &labels))) {
                                if (images != NULL) free(images);
                                if (labels != NULL) free(labels);
                                return err;
                        }
                }
        }

        // ---
        // forward pass
        {
                NO_ERR(vmExec(vm, OP_MATMUL, NULL, h1, x, w1));
                NO_ERR(vmExec(vm, OP_ADD, NULL, h1b, h1, b1));
                NO_ERR(vmExec(vm, OP_MAX, NULL, z1, h1b, z));
                NO_ERR(vmExec(vm, OP_MATMUL, NULL, h2, z1, w2));
                NO_ERR(vmExec(vm, OP_ADD, NULL, h2b, h2, b2));
                NO_ERR(vmExec(vm, OP_MAX, NULL, z2, h2b, z));
                NO_ERR(vmExec(vm, OP_MATMUL, NULL, o, z2, w3));
                NO_ERR(vmExec(vm, OP_LS_SCEL, NULL, loss, y, o));
        }

cleanup:
        if (images != NULL) free(images);
        if (labels != NULL) free(labels);
        srng64Free(seed);
        vmFree(vm);
        return err;
}

// impl
error_t readMnistData(unsigned char** images, unsigned char** labels)
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

void prepareFakeData(struct srng64_t* seed, float32_t* x_data, size_t x_size,
                     float32_t* y_data, size_t y_size)
{
        srng64StdNormalF(seed, x_size, x_data);
        srng64StdNormalF(seed, y_size, y_data);
}

error_t initModelWeight(struct vm_t* vm, struct srng64_t* seed,
                        struct opopt_t* opt, int w)
{
        struct srng64_t* weight_seed = srng64Split(seed);
        opt->rng_seed                = weight_seed;
        error_t err = vmExec(vm, OP_RNG, opt, w, VM_UNUSED, VM_UNUSED);
        srng64Free(weight_seed);

        return err;
}
