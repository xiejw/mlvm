#include "vm.h"

#include <stdio.h>

#include "opcode.h"

#define STACK_INIT_SIZE 256

#define DEBUG_PRINT(x) printf(x)

error_t handleOpCode(struct vm_t* vm, code_t** pc, enum opcode_t op);

// -----------------------------------------------------------------------------
// implementation.
// -----------------------------------------------------------------------------

struct vm_t* vmNew(void)
{
        struct vm_t* vm     = malloc(sizeof(struct vm_t));
        vm->size_used       = 0;
        struct obj_t* stack = malloc(STACK_INIT_SIZE * sizeof(struct obj_t));
        vm->stack           = stack;
        vm->base            = stack;
        vm->top             = stack;
        return vm;
}

void vmFree(struct vm_t* vm)
{
        if (vm == NULL) return;

        free(vm->stack);
        free(vm);
        // objGC();
}

error_t vmLaunch(struct vm_t* vm, vec_t(code_t) code)
{
        enum opcode_t op;
        code_t*       pc = code;

        while (1) {
                op = (enum opcode_t) * pc++;
                if (op != OP_HALT) {
                        if (handleOpCode(vm, &pc, op))
                                return errEmitNote(
                                    "unexpected op handing error in vm.");
                } else {
                        DEBUG_PRINT("vm halt\n");
                        return OK;
                }
        }
        return OK;
}

float vmComsumedSizeInMB(struct vm_t* vm)
{
        return (float)(((double)vm->size_used) / 1024 / 1024);
}

// internal
error_t handleOpCode(struct vm_t* vm, code_t** pc, enum opcode_t op)
{
        switch (op) {
        case OP_PUSHBYTE:
                (vm->top++)->value.i = *(*pc)++;
                break;
        default:
                return errNew("unsupported opcode: %d", op);
        }

        return OK;
}
