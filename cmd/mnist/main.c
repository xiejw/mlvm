#include <stdio.h>

#include "adt/sds.h"
#include "rng/srng64.h"
#include "rng/srng64_normal.h"

#include "vm.h"

#include "mnist.c"

static error_t readMnistData(unsigned char** images, unsigned char** labels);
static void prepareFakeData(struct srng64_t* seed, float32_t* x_data, size_t);

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

int main()
{
        error_t          err    = OK;
        unsigned char*   images = NULL;
        unsigned char*   labels = NULL;
        struct srng64_t* seed   = srng64New(123);

        if (!FAKE_DATA) {
                printf("reading real minis data.");
                if ((err = readMnistData(&images, &labels))) {
                        if (images != NULL) free(images);
                        if (labels != NULL) free(labels);
                        return err;
                }
        }

        // is = IMAGE_SIZE
        // ls = LABEL_SIZE
        //
        // x[bs, is]
        // y[bs]
        //
        // forward pass
        //
        //   z[1] = zeros([1])
        //   h1[bs, h1] = matmul(x[bs, is] w1[is, h1])
        //   z1[bs, h1] = max(h1[bs, h1], z[1])
        //   h2[bs, h2] = matmul(z1[bs, h1] w2[h1, h2])
        //   z2[bs, h2] = max(h2[bs, h2], z[1])
        //   o [bs, ls] = matmul(z2[bs, h2] w3[h2, ls])
        //   l [bs, ls] = one_hot(y[bs], depth=ls)
        //   loss[1] = softmax_cross_entropy_with_logits(l[bs, ls], o[bs, ls])

        const int bs = 32;
        const int is = IMAGE_SIZE;
        // const int ls   = LABEL_SIZE;
        // const int h1_s = 64;
        // const int h2_s = 64;

        // ---
        // defines vm with some shapes.

        struct vm_t* vm = vmNew();
        // struct shape_t* sp_scalar = spNew(1, (int[]){1});
        struct shape_t* sp_x = spNew(2, (int[]){bs, is});
        // struct shape_t* sp_y      = spNew(1, (int[]){bs});
        // struct shape_t* sp_w1     = spNew(2, (int[]){is, h1_s});
        // struct shape_t* sp_h1     = spNew(2, (int[]){bs, h1_s});
        // struct shape_t* sp_w2     = spNew(2, (int[]){h1_s, h2_s});
        // struct shape_t* sp_h2     = spNew(2, (int[]){bs, h2_s});
        // struct shape_t* sp_o      = spNew(2, (int[]){bs, ls});
        // struct shape_t* sp_l      = spNew(2, (int[]){bs, ls});
        // struct opopt_t  opt;

        int x = vmTensorNew(vm, F32, sp_x);
        // int y    = vmTensorNew(vm, F32, sp_y);
        // int z    = vmTensorNew(vm, F32, sp_scalar);
        // int w1   = vmTensorNew(vm, F32, sp_w1);
        // int h1   = vmTensorNew(vm, F32, sp_h1);
        // int z1   = vmTensorNew(vm, F32, sp_h1);
        // int h2   = vmTensorNew(vm, F32, sp_h2);
        // int z2   = vmTensorNew(vm, F32, sp_h2);
        // int o    = vmTensorNew(vm, F32, sp_o);
        // int l    = vmTensorNew(vm, F32, sp_l);
        // int loss = vmTensorNew(vm, F32, sp_scalar);

        float32_t* x_data;

        NO_ERR(vmTensorData(vm, x, (void**)&x_data));

        if (FAKE_DATA) {
                printf("generating fake minis data.");
                prepareFakeData(seed, x_data, /*size=*/ sp_x->size);
        }

cleanup:
        spFreeAll();
        if (images != NULL) free(images);
        if (labels != NULL) free(labels);
        srng64Free(seed);
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

void prepareFakeData(struct srng64_t* seed, float32_t* x_data, size_t size)
{
        srng64StdNormalF(seed, size, x_data);
}
