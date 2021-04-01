#include <stdio.h>

#include "adt/sds.h"
#include "vm.h"

int main()
{
        error_t err = OK;
        sds_t   s   = sdsEmpty();
        printf("hello mlvm\n");
        struct vm_t*    vm     = vmNew();
        struct shape_t* r2_2x3 = spNew(2, (int[]){2, 3});

        int t1 = vmTensorNew(vm, F32, r2_2x3);
        int t2 = vmTensorNew(vm, F32, r2_2x3);
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
        sdsFree(s);
        vmFree(vm);
        return err;
}
