#include "stdio.h"

#include "adt/sds.h"
#include "vm/object.h"
#include "vm/opcode.h"
#include "vm/vm.h"

#define CHECK(x, msg)                                              \
        do {                                                       \
                if (x) {                                           \
                        errDump(msg);                              \
                        errFree();                                 \
                        return errFatalAndExit("unexpected err."); \
                }                                                  \
        } while (0)

int main()
{
        struct vm_t* vm                     = vmNew();
        vec_t(code_t) code                  = vecNew();
        vec_t(struct obj_tensor_t*) outputs = vecNew();

        vm_handle_t handle = vmAllocTensor(vm, 1, (int[]){1});
        assert(handle >= 0);

        CHECK(opMake(&code, OP_PUSHBYTE, handle), "program op error");
        CHECK(opMake(&code, OP_LOADGLOBAL), "program op error");
        CHECK(opMake(&code, OP_HALT), "program op error");

        CHECK(vmLaunch(vm, code, &outputs), "vm execution error");

        sds_t s = sdsEmpty();
        for (int i = 0; i < vecSize(outputs); i++) {
                struct obj_tensor_t* t = outputs[i];
                objTensorDump(t, &s);
                printf("output %d has rank %d and value: %s\n", i, t->rank, s);

                objTensorFree(t);
        }
        sdsFree(s);
        vmFree(vm);
        vecFree(code);
        return 0;
}
