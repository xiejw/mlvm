#include "stack.h"

#include <stdio.h>
#include <string.h>

#include "adt/vec.h"
#include "object.h"
#include "opcode.h"

#define STACK_INIT_SIZE 256

static struct obj_t *base  = NULL;
static struct obj_t *top   = NULL;
static struct obj_t *stack = NULL;

static code_t *pc = NULL;

#define DEBUG_PRINT(x) printf(x)

error_t handleOpCode(enum opcode_t op);

void vmInit()
{
        if (stack != NULL) free(stack);
        stack = malloc(STACK_INIT_SIZE * sizeof(struct obj_t));
        base  = stack;
        top   = stack;
}

void vmFree() {
        if (stack != NULL) {
                free(stack);
                stack = NULL;
        }
        objGC();
}

error_t vmExec(vec_t(code_t) code)
{
        enum opcode_t op;
        pc = code;

        while (1) {
                op = (enum opcode_t) * pc++;
                if (op != OP_HALT) {
                        if (handleOpCode(op))
                                return errEmitNote(
                                    "unexpected op handing error in vm.");
                } else {
                        DEBUG_PRINT("vm halt\n");
                        return OK;
                }
        }
}

error_t handleOpCode(enum opcode_t op)
{
        switch (op) {
                case OP_PUSHBYTE:
                        (top++)->value.i = *pc++;
                        break;
                default:
                        return errNew("unsupported opcode: %d", op);
        }

        return OK;
}
