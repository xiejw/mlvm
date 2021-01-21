#include "stdio.h"

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
        struct vm_t* vm    = vmNew();
        vec_t(code_t) code = vecNew();

        CHECK(opMake(&code, OP_PUSHBYTE, 1), "program op error");
        CHECK(opMake(&code, OP_HALT), "program op error");

        CHECK(vmLaunch(vm, code), "vm execution error");

        vmFree(vm);
        vecFree(code);
        return 0;
}
