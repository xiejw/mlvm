#include "stdio.h"

#include "vm/opcode.h"
#include "vm/stack.h"

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
        vmInit();
        vec_t(code_t) code = vecNew();

        CHECK(opMake(&code, OP_PUSHBYTE, 1), "program op error");
        CHECK(opMake(&code, OP_HALT), "program op error");

        CHECK(vmExec(code), "vm execution error");

        vmFree();
        vecFree(code);
        return 0;
}
