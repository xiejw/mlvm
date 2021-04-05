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

int main()
{
        error_t err = OK;
        sds_t   s   = sdsEmpty();

        struct vm_t*     vm     = vmNew();
        struct shape_t*  r2_2x3 = spNew(2, (int[]){2, 3});
        struct srng64_t* rng    = srng64New(123);
        struct opopt_t   opt;

        printf("hello mlvm\n");

        int t1       = vmTensorNew(vm, F32, r2_2x3);
        opt.mode     = 0;  // normal.
        opt.rng_seed = rng;
        NO_ERR(vmExec(vm, OP_RNG, &opt, t1, VM_UNUSED, VM_UNUSED));

        int t2 = vmTensorNew(vm, F32, r2_2x3);
        NO_ERR(vmExec(vm, OP_RNG, &opt, t2, VM_UNUSED, VM_UNUSED));

        sdsCatPrintf(&s, "t1: ");
        vmTensorDump(&s, vm, t1);
        sdsCatPrintf(&s, "\n");

        sdsCatPrintf(&s, "t2: ");
        vmTensorDump(&s, vm, t2);
        sdsCatPrintf(&s, "\n");

        NO_ERR(vmExec(vm, OP_ADD, NULL, t1, t1, t2));

        sdsCatPrintf(&s, "ds: ");
        vmTensorDump(&s, vm, t1);
        sdsCatPrintf(&s, "\n");

        NO_ERR(vmExec(vm, OP_MUL, NULL, t1, t1, t2));

        sdsCatPrintf(&s, "ds: ");
        vmTensorDump(&s, vm, t1);
        sdsCatPrintf(&s, "\n");

        printf("%s\n", s);

cleanup:
        spDecRef(r2_2x3);
        srng64Free(rng);
        sdsFree(s);
        vmFree(vm);
        return err;
}
