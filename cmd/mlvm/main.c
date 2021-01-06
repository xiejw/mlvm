#include "stdio.h"

#include "vm/stack.h"

int main()
{
        vec_t(opcode_t) code = vecNew();
        vecPushBack(code, OP_HALT);
        if (vmExec(code)) {
                errDump("vm execution error:");
                return errFatalAndExit("unexpected err.");
        }

        vecFree(code);
        return 0;
}
