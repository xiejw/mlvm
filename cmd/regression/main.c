#include <stdio.h>

#include "adt/sds.h"
#include "rng/srng64.h"
#include "rng/srng64_normal.h"

#include "vm.h"

#define NO_ERR(e) NO_ERR_IMPL_(e, __FILE__, __LINE__)

#define NO_ERR_IMPL_(e, f, l)                                           \
        if (e) {                                                        \
                err = e;                                                \
                errDump("failed to exec op @ file %s line %d\n", f, l); \
                goto cleanup;                                           \
        }

#define SDS_CAT_PRINTF(prefix, t, suffix) \
        sdsCatPrintf(&s, prefix);         \
        vmTensorDump(&s, vm, t);          \
        sdsCatPrintf(&s, suffix);

// helper to generate next input
void new_input(struct srng64_t* seed, size_t size, _mut_ float32_t* data,
               float32_t* y, float32_t* w);

int main()
{
        error_t err = OK;
        sds_t   s   = sdsEmpty();

        printf("Linear Regression\n");

        // ---
        // defines vm with some shapes.

        struct vm_t*    vm        = vmNew();
        struct shape_t* sp_weight = spNew(1, (int[]){6});
        struct shape_t* r1_1      = spNew(1, (int[]){1});
        struct opopt_t  opt;

        // ---
        // prepares the seeds, one for model and one for input.

        struct srng64_t* seed           = srng64New(123);
        struct srng64_t* seed_for_input = srng64Split(seed);
        struct srng64_t* rng;  // free after each use.

        // ---
        // allocates the tensors for model, input and output.

        int w_target = vmTensorNew(vm, F32, sp_weight);
        int w        = vmTensorNew(vm, F32, sp_weight);
        int x        = vmTensorNew(vm, F32, sp_weight);
        int y        = vmTensorNew(vm, F32, r1_1);

        // ---
        // obtains the pointers to the real data.

        float32_t* x_data;
        float32_t* y_data;
        float32_t* w_data;

        NO_ERR(vmTensorData(vm, x, (void**)&x_data));
        NO_ERR(vmTensorData(vm, w_target, (void**)&w_data));
        NO_ERR(vmTensorData(vm, y, (void**)&y_data));

        // ---
        // initializes weight for the model (target).
        rng          = srng64Split(seed);
        opt.mode     = 0;  // normal.
        opt.rng_seed = rng;
        NO_ERR(vmExec(vm, OP_RNG, &opt, w_target, VM_UNUSED, VM_UNUSED));
        srng64Free(rng);

        SDS_CAT_PRINTF("\ttarget  weight: ", w_target, "\n");

        // ---
        // initializes weight for the model (about to learn).
        rng          = srng64Split(seed);
        opt.rng_seed = rng;
        NO_ERR(vmExec(vm, OP_RNG, &opt, w, VM_UNUSED, VM_UNUSED));
        srng64Free(rng);

        SDS_CAT_PRINTF("\tinitial weight: ", w, "\n");

        // --- formula
        //   forward pass
        //      z[w] = x[w] * w[w]
        //      rz[] = reduce_sum(z[w])
        //      l[] = rz[] - y[]
        //      l2[] = l[] * l[]
        //      loss[] = reduce_sum(l2[])
        //   backward pass
        //      d_l2[] = ones_like(l2[])
        //      d_l[] = 2 * l[] * d_l2[]
        //      d_rz[] = d_l[]
        //      d_z[w] = ones_like(x[w]) * d_rz[]
        //      d_w[w] = d_z[w] * x[w]
        //
        //   optimizer
        //      d_w[w] = lr[] * d_w[w]  // reuse: update
        //      w[w] = w[w] - d_w[w]
        //
        //   backward pass + optimizer (simplified)
        //      d_rz[] = (2*lr) * l[]
        //      d_w[w] = d_rz[] * x[w]
        //      w[w] = w[w] - d_w[w]

        int z    = vmTensorNew(vm, F32, sp_weight);
        int rz   = vmTensorNew(vm, F32, r1_1);
        int l    = vmTensorNew(vm, F32, r1_1);
        int l2   = vmTensorNew(vm, F32, r1_1);
        int loss = l2;
        int d_rz = vmTensorNew(vm, F32, r1_1);
        int d_w  = vmTensorNew(vm, F32, sp_weight);

        for (size_t i = 0; i < 100; i++) {
                sdsCatPrintf(&s, "\niteration %d\n", i);

                // ---
                // prepare input sample: x,y.
                new_input(seed_for_input, sp_weight->size, x_data, y_data,
                          w_data);
                SDS_CAT_PRINTF("\tinput: ", x, "\n");
                SDS_CAT_PRINTF("\ty: ", y, "\n");

                // forward pass.
                NO_ERR(vmExec(vm, OP_MUL, NULL, z, x, w));
                NO_ERR(vmExec(vm, OP_REDUCE, &opt, rz, z, VM_UNUSED));
                NO_ERR(vmExec(vm, OP_MINUS, NULL, l, rz, y));
                NO_ERR(vmExec(vm, OP_MUL, NULL, l2, l, l));
                SDS_CAT_PRINTF("\tloss : ", loss, "\n");

                // backward pass
                opt.f = 2 * 0.05;  // 2 * learning_rate
                NO_ERR(vmExec(vm, OP_MUL, &opt, d_rz, l, VM_UNUSED));
                NO_ERR(vmExec(vm, OP_MUL, NULL, d_w, x,
                              d_rz));  // d_rz must be t2.
                NO_ERR(vmExec(vm, OP_MINUS, NULL, w, w, d_w));
                SDS_CAT_PRINTF("\tgrad : ", d_w, "\n");
                SDS_CAT_PRINTF("\tnew_w: ", w, "\n");
                SDS_CAT_PRINTF("\ttgt w: ", w_target, "\n");
                printf("%s\n", s);
                sdsClear(s);
        }

cleanup:
        spDecRef(sp_weight);
        spDecRef(r1_1);
        srng64Free(seed_for_input);
        srng64Free(seed);
        sdsFree(s);
        vmFree(vm);
        return err;
}

// internal

void new_input(struct srng64_t* seed, size_t size, _mut_ float32_t* data,
               float32_t* y, float32_t* w)
{
        // x
        srng64StdNormalF(seed, size, data);

        // y
        float32_t local_y = 0;
        for (size_t i = 0; i < size; i++) {
                local_y += w[i] * data[i];
        }
        *y = local_y;
}