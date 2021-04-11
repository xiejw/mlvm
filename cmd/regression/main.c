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

        struct vm_t*    vm        = vmNew();
        struct shape_t* sp_weight = spNew(1, (int[]){6});
        struct shape_t* r1_1      = spNew(1, (int[]){1});
        struct opopt_t  opt;

        struct srng64_t* seed           = srng64New(123);
        struct srng64_t* seed_for_input = srng64Split(seed);
        struct srng64_t* rng;  // free after each use.

        float32_t* x;
        float32_t  y;
        float32_t* w_data;

        printf("Linear Regression\n");

        int w_target = vmTensorNew(vm, F32, sp_weight);
        int w_learn  = vmTensorNew(vm, F32, sp_weight);
        int input    = vmTensorNew(vm, F32, sp_weight);

        NO_ERR(vmTensorData(vm, input, (void**)&x));
        NO_ERR(vmTensorData(vm, w_target, (void**)&w_data));

        // ---
        // initial weight for the model (target).
        opt.mode     = 0;  // normal.
        rng          = srng64Split(seed);
        opt.rng_seed = rng;
        NO_ERR(vmExec(vm, OP_RNG, &opt, w_target, VM_UNUSED, VM_UNUSED));
        srng64Free(rng);

        SDS_CAT_PRINTF("\ttarget  weight: ", w_target, "\n");

        // ---
        // initial weight for the model (about to learn).
        rng          = srng64Split(seed);
        opt.rng_seed = rng;
        NO_ERR(vmExec(vm, OP_RNG, &opt, w_learn, VM_UNUSED, VM_UNUSED));
        srng64Free(rng);

        SDS_CAT_PRINTF("\tinitial weight: ", w_learn, "\n");

        // --
        // prepare input sample: x,y.
        new_input(seed_for_input, sp_weight->size, x, &y, w_data);
        SDS_CAT_PRINTF("\tinput: ", input, "\n");
        sdsCatPrintf(&s, "\ty: %f\n", y);

        // jNO_ERR(vmExec(vm, OP_ADD, NULL, t1, t1, t2));
        // jSDS_CAT_PRINTF("t1 <- t1 + t2\n\tt1: ", t1, "\n");

        // jNO_ERR(vmExec(vm, OP_MUL, NULL, t1, t1, t2));
        // jSDS_CAT_PRINTF("t1 <- t1 * t2\n\tt1: ", t1, "\n");

        // jNO_ERR(vmExec(vm, OP_REDUCE, &opt, t3, t1, VM_UNUSED));
        // jSDS_CAT_PRINTF("t3 <- reduce(t1)\n\tt3: ", t3, "\n");

        printf("%s\n", s);

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
