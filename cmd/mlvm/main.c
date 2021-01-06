#include "stdio.h"

#include "vm/opcode.h"
#include "vm/stack.h"

#define ADD_OP(x)                                                  \
        do {                                                       \
                if (x) {                                           \
                        errDump("program op error");               \
                        return errFatalAndExit("unexpected err."); \
                }                                                  \
        } while (0)

int main()
{
        vec_t(code_t) code = vecNew();
        ADD_OP(opMake(&code, OP_PUSHBYTE, 1));
        ADD_OP(opMake(&code, OP_HALT));

        vmInit();

        if (vmExec(code)) {
                errDump("vm execution error");
                return errFatalAndExit("unexpected err.");
        }

        vmFree();

        vecFree(code);
        return 0;
}
