#include <stdio.h>

#include "vm.h"

int main()
{
        error_t err = OK;
        printf("hello mlvm\n");
        struct vm_t*    vm     = vmNew();
        struct shape_t* r2_2x3 = spNew(2, (int[]){2, 3});

        int t1 = vmTensorNew(vm, F32, r2_2x3);
        int t2 = vmTensorNew(vm, F32, r2_2x3);

        err = vmExec(vm, OP_ADD, NULL, t1, t1, t2);
        if (err) {
                errDump("failed to exec op");
                goto cleanup;
        }

        printf("succeed.\n");

cleanup:
        spDecRef(r2_2x3);
        vmFree(vm);
        return err;
}
