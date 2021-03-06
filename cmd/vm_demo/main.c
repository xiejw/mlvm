#include <stdio.h>

// eva
#include "adt/sds.h"
#include "rng/srng64.h"

// mlvm
#include "vm.h"

// cmd
#include "../helpers.h"  // NE, S_PRINTF macros

int
main()
{
        error_t err = OK;

        struct vm_t     *vm     = vmNew();
        struct shape_t  *r2_2x3 = spNew(2, (int[]){2, 3});
        struct shape_t  *r1_1   = spNew(1, (int[]){1});
        struct srng64_t *rng    = srng64New(123);
        struct opopt_t   opt;

        sds_t s = sdsEmpty();

        // Allocates some tensors.
        int t1 = vmTensorNew(vm, F32, r2_2x3);  // rank 2, with shape <2, 3>
        int t2 = vmTensorNew(vm, F32, r2_2x3);  // rank 2, with shape <2, 3>
        int t3 = vmTensorNew(vm, F32, r1_1);    // scalar

        {
                // Fills t1 and t2 with random numbers with the same seed.
                opt.mode = OPT_RNG_STD_NORMAL | OPT_MODE_R_BIT;
                opt.r    = *(struct rng64_t *)rng;
                NE(vmExec(vm, OP_RNG, &opt, t1, -1, -1));
                NE(vmExec(vm, OP_RNG, &opt, t2, -1, -1));

                S_PRINTF("Init values:\n\tt1: ", t1, "\n");
                S_PRINTF("\tt2: ", t2, "\n");
        }

        {
                // Performs mutation t1 = t1 + t2.
                NE(vmExec(vm, OP_ADD, NULL, t1, t1, t2));
                S_PRINTF("t1 <- t1 + t2\n\tt1: ", t1, "\n");
        }

        {
                // Performs mutation t1 = t1 * t2.
                NE(vmExec(vm, OP_MUL, NULL, t1, t1, t2));
                S_PRINTF("t1 <- t1 * t2\n\tt1: ", t1, "\n");
        }

        {
                // Performs reduction.
                OPT_SET_REDUCTION_SUM(opt, 0);
                NE(vmExec(vm, OP_REDUCE, &opt, t3, t1, -1));
                S_PRINTF("t3 <- reduce(t1)\n\tt3: ", t3, "\n");
        }

        printf("Hello MLVM\n\n%s\n", s);

cleanup:
        spDecRef(r2_2x3);
        spDecRef(r1_1);
        srng64Free(rng);
        sdsFree(s);
        vmFree(vm);
        return err;
}
