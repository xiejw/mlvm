#include <stdio.h>

// eva
#include "adt/sds.h"
#include "adt/vec.h"
#include "rng/srng64.h"
#include "rng/srng64_normal.h"

// mlvm
#include "vm.h"

// cmd
#include "../helpers.h"
#include "mnist.h"

static error_t initTensorWRng(struct vm_t*, struct srng64_t*, int w);
static error_t initTensorWZeros(struct vm_t*, int w);
static error_t prepareData(struct srng64_t* seed, float32_t* x_data,
                           size_t x_size, float32_t* y_data, size_t y_size);

#define IMAGE_SIZE (28 * 28)
#define LABEL_SIZE (10)

#define FAKE_DATA 1

static unsigned char* images = NULL;
static unsigned char* labels = NULL;

static vec_t(int) weights = vecNew();
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
                //   d_z2[bs, h2] = matmul(d_o[bs, ls], w3[h2, ls], trans_b)
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
                //
                // optimizer
                //     for all weights w_i
                //       d_wi[*] = lr[] * d_wi[*]
                //       wi[*]   = wi[*] - d_wi[*]
        }

        const int bs   = 32;
        const int is   = IMAGE_SIZE;
        const int ls   = LABEL_SIZE;
        const int h1_s = 64;
        const int h2_s = 64;

        // ---
        // defines vm with some shapes.
        struct vm_t* vm = vmNew();

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

        int arg_y = vmTensorNew(vm, F32, sp_l);
        int arg_o = vmTensorNew(vm, F32, sp_l);
        int same  = vmTensorNew(vm, F32, sp_l);
        int count = vmTensorNew(vm, F32, sp_scalar);

        int d_o  = vmTensorNew(vm, F32, sp_o);
        int d_w3 = vmTensorNew(vm, F32, sp_w3);
        int d_z2 = vmTensorNew(vm, F32, sp_h2);

        int state_0 = vmTensorNew(vm, F32, sp_h2);
        int d_h2b   = vmTensorNew(vm, F32, sp_h2);
        int d_h2    = d_h2b;
        int d_b2    = vmTensorNew(vm, F32, sp_b2);
        int d_w2    = vmTensorNew(vm, F32, sp_w2);
        int d_z1    = vmTensorNew(vm, F32, sp_h1);

        int state_1 = vmTensorNew(vm, F32, sp_h1);
        int d_h1b   = vmTensorNew(vm, F32, sp_h1);
        int d_h1    = d_h1b;
        int d_b1    = vmTensorNew(vm, F32, sp_b1);
        int d_w1    = vmTensorNew(vm, F32, sp_w1);

        // ---
        // init weights
        {
                printf("init model weights.\n");
                NE(initTensorWRng(vm, seed, w1));
                NE(initTensorWZeros(vm, b1));
                NE(initTensorWRng(vm, seed, w2));
                NE(initTensorWZeros(vm, b2));
                NE(initTensorWRng(vm, seed, w3));
                vecPushBack(weights, w1);
                vecPushBack(weights, b1);
                vecPushBack(weights, w2);
                vecPushBack(weights, b2);
                vecPushBack(weights, w3);
                S_PRINTF("w1: ", w1, "\n");
                S_PRINTF("b1: ", b1, "\n");
                S_PRINTF("w2: ", w2, "\n");
                S_PRINTF("b2: ", b2, "\n");
                S_PRINTF("w2: ", w3, "\n");
                printf("%s\n", s);
                sdsClear(s);
        }
        NE(initTensorWZeros(vm, z));
        NE(initTensorWZeros(vm, count));
        // NE(initTensorWZeros(vm, total));

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
        // loop
        {
                const struct oparg_t prog[] = {
                    // clang-format off
// the first matmul
{OP_MATMUL,  h1,      x,     w1,      0},
{OP_ADD,     h1b,     h1,    b1,      0},
{OP_MAX,     z1,      h1b,   z,       0},
// the second matmul
{OP_MATMUL,  h2,      z1,    w2,      0},
{OP_ADD,     h2b,     h2,    b2,      0},
{OP_MAX,     z2,      h2b,   z,       0},
// the linear layer
{OP_MATMUL,  o,       z2,    w3,      0},
{OP_LS_SCEL, l,       y,     o,       1, {.mode=OPT_MODE_I_BIT,       .i=d_o  }},
{OP_REDUCE,  loss,    l,     -1,      1, {.mode=0|OPT_MODE_I_BIT,     .i=0    }},

// backprop for the linear layer
{OP_MATMUL,  d_w3,    z2,    d_o,     1, {.mode=OPT_MATMUL_TRANS_LHS          }},
{OP_MATMUL,  d_z2,    d_o,   w3,      1, {.mode=OPT_MATMUL_TRANS_RHS          }},
// backprop for the second matmul
{OP_CMPL,    state_0, h2b,   z,       0},
{OP_MUL,     d_h2b,   d_z2,  state_0, 0},
{OP_REDUCE,  d_b2,    d_h2b, -1,      1, {.mode=0|OPT_MODE_I_BIT,     .i=1    }},
{OP_MATMUL,  d_w2,    z1,    d_h2,    1, {.mode=OPT_MATMUL_TRANS_LHS          }},
{OP_MATMUL,  d_z1,    d_h2,  w2,      1, {.mode=OPT_MATMUL_TRANS_RHS          }},
// backprop for the first matmul
{OP_CMPL,    state_1, h1b,   z,       0},
{OP_MUL,     d_h1b,   d_z1,  state_1, 0},
{OP_REDUCE,  d_b1,    d_h1b, -1,      1, {.mode=0|OPT_MODE_I_BIT,     .i=1    }},
{OP_MATMUL,  d_w1,    x,     d_h1,    1, {.mode=OPT_MATMUL_TRANS_LHS          }},

// optimizer
{OP_MUL,     d_w1,    d_w1,  -1,      1, {.mode=0|OPT_MODE_F_BIT,     .f=.001  }},
{OP_MUL,     d_b1,    d_b1,  -1,      1, {.mode=0|OPT_MODE_F_BIT,     .f=.001  }},
{OP_MUL,     d_w2,    d_w2,  -1,      1, {.mode=0|OPT_MODE_F_BIT,     .f=.001  }},
{OP_MUL,     d_b2,    d_b2,  -1,      1, {.mode=0|OPT_MODE_F_BIT,     .f=.001  }},
{OP_MUL,     d_w3,    d_w3,  -1,      1, {.mode=0|OPT_MODE_F_BIT,     .f=.001  }},

{OP_MINUS,   w1,      w1,    d_w1,    0},
{OP_MINUS,   b1,      b1,    d_b1,    0},
{OP_MINUS,   w2,      w2,    d_w2,    0},
{OP_MINUS,   b2,      b2,    d_b2,    0},
{OP_MINUS,   w3,      w3,    d_w3,    0},
                    // clang-format on
                };

                for (int i = 0; i < 100; i++) {
                        NE(vmBatch(vm, sizeof(prog) / sizeof(struct oparg_t),
                                   prog));

                        S_PRINTF("z: ", z, "\n");
                        S_PRINTF("x: ", x, "\n");
                        S_PRINTF("w1: ", w1, "\n");
                        S_PRINTF("h1: ", h1, "\n");
                        S_PRINTF("b1: ", b1, "\n");
                        S_PRINTF("h1b: ", h1b, "\n");
                        S_PRINTF("z1: ", z1, "\n");
                        S_PRINTF("logits: ", o, "\n");
                        S_PRINTF("labels: ", y, "\n");
                        S_PRINTF("loss after scel: ", l, "\n");
                        S_PRINTF("loss: ", loss, "\n");
                        S_PRINTF("grad d_o: ", d_o, "\n");
                        S_PRINTF("d_w1: ", d_w1, "\n");
                        printf("%s\n", s);
                        sdsClear(s);

                        NE(vmExec(vm, OP_ARGMAX, NULL, arg_y, y, -1));
                        NE(vmExec(vm, OP_ARGMAX, NULL, arg_o, o, -1));
                        NE(vmExec(vm, OP_EQ, NULL, same, arg_y, arg_o));
                        struct opopt_t opt;
                        OPT_SET_REDUCTION_SUM(opt, 0);
                        NE(vmExec(vm, OP_REDUCE, &opt, count, same, -1));
                        S_PRINTF("count: ", count, "\n");
                        printf("%s\n", s);
                        sdsClear(s);
                }
        }

        S_PRINTF("w1: ", w1, "\n");
        S_PRINTF("b1: ", b1, "\n");
        S_PRINTF("w2: ", w2, "\n");
        S_PRINTF("b2: ", b2, "\n");
        S_PRINTF("w2: ", w3, "\n");
        printf("%s\n", s);
        sdsClear(s);

cleanup:
        if (images != NULL) free(images);
        if (labels != NULL) free(labels);
        free(seed);
        vmFree(vm);
        vecFree(weights);
        sdsFree(s);
        return err;
}

// impl
static error_t
prepareMnistData(unsigned char** images, unsigned char** labels)
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

#include <stdio.h>

static void
prepareFakeData(struct srng64_t* seed, float32_t* x_data, size_t x_size,
                float32_t* y_data, size_t y_size)
{
        rng64StdNormalF((struct rng64_t*)seed, x_size, x_data);
        size_t bs = y_size / LABEL_SIZE;

        struct rng64_t* rng = (struct rng64_t*)seed;
        for (size_t i = 0; i < bs; i++) {
                int target = rng64NextUint64(rng) % LABEL_SIZE;
                printf("tgt: %d\n", target);
                for (size_t j = 0; j < LABEL_SIZE; j++) {
                        if (j == target)
                                y_data[i * LABEL_SIZE + j] = 1;
                        else
                                y_data[i * LABEL_SIZE + j] = 0;
                }
        }
}

error_t
prepareData(struct srng64_t* seed, float32_t* x_data, size_t x_size,
            float32_t* y_data, size_t y_size)
{
        if (FAKE_DATA) {
                printf("generating fake minis data.\n");
                prepareFakeData(seed, x_data, x_size, y_data, y_size);
                return OK;
        } else {
                error_t err;
                printf("reading real minis data.\n");
                if ((err = prepareMnistData(&images, &labels))) {
                        if (images != NULL) free(images);
                        if (labels != NULL) free(labels);
                        return err;
                }
                return OK;
        }
}

error_t
initTensorWRng(struct vm_t* vm, struct srng64_t* seed, int w)
{
        struct srng64_t* rng = srng64Split(seed);

        struct opopt_t opt;
        opt.mode    = OPT_RNG_STD_NORMAL | OPT_MODE_R_BIT;
        opt.r       = *(struct rng64_t*)rng;
        error_t err = vmExec(vm, OP_RNG, &opt, w, -1, -1);

        free(rng);
        return err;
}

error_t
initTensorWZeros(struct vm_t* vm, int w)
{
        return vmExec(vm, OP_FILL, NULL, w, -1, -1);
}
