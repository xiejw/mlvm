#include <stdio.h>

#include "adt/sds.h"
#include "rng/srng64.h"

#include "vm.h"

int main()
{
        error_t err = OK;
        sds_t   s   = sdsEmpty();

        printf("hello mlvm\n");
        struct srng64_t* rng    = srng64New(123);
        struct vm_t*     vm     = vmNew();
        struct opopt_t*  opt    = vmOptNew();
        struct shape_t*  r2_2x3 = spNew(2, (int[]){2, 3});

        int t1        = vmTensorNew(vm, F32, r2_2x3);
        opt->mode     = 0;  // normal.
        opt->rng_seed = rng;
        err           = vmExec(vm, OP_RNG, opt, t1, -1, -1);
        if (err) {
                errDump("failed to exec op");
                goto cleanup;
        }

        int t2 = vmTensorNew(vm, F32, r2_2x3);
        err    = vmExec(vm, OP_RNG, opt, t2, -1, -1);
        if (err) {
                errDump("failed to exec op");
                goto cleanup;
        }

        sdsCatPrintf(&s, "t1: ");
        vmTensorDump(&s, vm, t1);
        sdsCatPrintf(&s, "\n");

        sdsCatPrintf(&s, "t2: ");
        vmTensorDump(&s, vm, t2);
        sdsCatPrintf(&s, "\n");

        err = vmExec(vm, OP_ADD, NULL, t1, t1, t2);
        if (err) {
                errDump("failed to exec op");
                goto cleanup;
        }

        sdsCatPrintf(&s, "result: ");
        vmTensorDump(&s, vm, t1);

        printf("%s\n", s);

cleanup:
        spDecRef(r2_2x3);
        vmOptDecRef(opt);
        srng64Free(rng);
        sdsFree(s);
        vmFree(vm);
        return err;
}
