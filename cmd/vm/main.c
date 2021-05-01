#include <stdio.h>

#include "adt/sds.h"
#include "rng/srng64.h"

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

int
main()
{
        error_t err = OK;
        sds_t   s   = sdsEmpty();

        struct vm_t*     vm     = vmNew();
        struct shape_t*  r2_2x3 = spNew(2, (int[]){2, 3});
        struct shape_t*  r1_1   = spNew(1, (int[]){1});
        struct srng64_t* rng    = srng64New(123);
        struct opopt_t   opt;

        printf("hello mlvm\ninit\n");

        int t1 = vmTensorNew(vm, F32, r2_2x3);
        int t2 = vmTensorNew(vm, F32, r2_2x3);
        int t3 = vmTensorNew(vm, F32, r1_1);

        opt.mode     = 0;  // normal.
        opt.rng_seed = rng;
        NO_ERR(vmExec(vm, OP_RNG, &opt, t1, VM_UNUSED, VM_UNUSED));
        NO_ERR(vmExec(vm, OP_RNG, &opt, t2, VM_UNUSED, VM_UNUSED));
        SDS_CAT_PRINTF("\tt1: ", t1, "\n");
        SDS_CAT_PRINTF("\tt2: ", t2, "\n");

        NO_ERR(vmExec(vm, OP_ADD, NULL, t1, t1, t2));
        SDS_CAT_PRINTF("t1 <- t1 + t2\n\tt1: ", t1, "\n");

        NO_ERR(vmExec(vm, OP_MUL, NULL, t1, t1, t2));
        SDS_CAT_PRINTF("t1 <- t1 * t2\n\tt1: ", t1, "\n");

        OPT_SET_REDUCTION_SUM(opt);
        NO_ERR(vmExec(vm, OP_REDUCE, &opt, t3, t1, VM_UNUSED));
        SDS_CAT_PRINTF("t3 <- reduce(t1)\n\tt3: ", t3, "\n");

        printf("%s\n", s);

cleanup:
        spDecRef(r2_2x3);
        spDecRef(r1_1);
        srng64Free(rng);
        sdsFree(s);
        vmFree(vm);
        return err;
}
