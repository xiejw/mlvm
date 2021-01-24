#include "vm.h"

#include <stdio.h>
#include <string.h>

#include "opcode.h"

#define STACK_INIT_SIZE 256
#define MAX_NUM_HANDLES 128

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
        vm->handles = malloc(MAX_NUM_HANDLES * sizeof(struct obj_tensor_t*));
        memset(vm->handles, 0, MAX_NUM_HANDLES * sizeof(struct obj_tensor_t*));
        return vm;
}

void vmFree(struct vm_t* vm)
{
        if (vm == NULL) return;

        for (int i = 0; i < MAX_NUM_HANDLES; i++) {
                objTensorFree(vm->handles[i]);
        }

        free(vm->stack);
        free(vm->handles);
        free(vm);
        // objGC();
}

float vmComsumedSizeInMiB(struct vm_t* vm)
{
        return (float)(((double)vm->size_used) / 1024 / 1024);
}

vm_handle_t vmAllocateTensor(struct vm_t* vm, int rank, int dims[])
{
        int next_handle = -1;
        for (int i = 0; i < MAX_NUM_HANDLES; i++) {
                if (vm->handles[i] == NULL) {
                        next_handle = i;
                        break;
                }
        }

        if (next_handle == -1) return next_handle;
        struct obj_tensor_t* t    = objTensorNew(rank, dims);
        size_t               size = t->size * sizeof(obj_float_t);
        t->owned                  = 1;
        t->buffer                 = malloc(size);
        vm->handles[next_handle]  = t;
        vm->size_used += size;
        return next_handle;
}

error_t vmDeallocateTensor(struct vm_t* vm, vm_handle_t i)
{
        struct obj_tensor_t* t = vm->handles[i];
        if (t == NULL) return errNew("VM does not have tensor handle %d", i);
        vm->size_used -= t->size * sizeof(obj_float_t);
        objTensorFree(t);
        vm->handles[i] = NULL;
        return OK;
}

error_t vmRead(struct vm_t* vm, vm_handle_t i, obj_float_t* dst)
{
        struct obj_tensor_t* t = vm->handles[i];
        if (t == NULL) return errNew("VM does not have tensor handle %d", i);
        memcpy(dst, t->buffer, t->size * sizeof(obj_float_t));
        return OK;
}

error_t vmWrite(struct vm_t* vm, vm_handle_t i, obj_float_t* src)
{
        struct obj_tensor_t* t = vm->handles[i];
        if (t == NULL) return errNew("VM does not have tensor handle %d", i);
        memcpy(t->buffer, src, t->size * sizeof(obj_float_t));
        return OK;
}

error_t vmLaunch(struct vm_t* vm, vec_t(code_t) code,
                 vec_t(struct obj_tensor_t*) * outputs)
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
                        goto handle_outputs;
                }
        }

handle_outputs:

        return OK;
}

// -----------------------------------------------------------------------------
// internal.
// -----------------------------------------------------------------------------

error_t handleOpCode(struct vm_t* vm, code_t** pc, enum opcode_t op)
{
        vm_handle_t   handle;
        struct obj_t* top;

        switch (op) {
        case OP_PUSHBYTE:
                (vm->top++)->value.i = *((*pc)++);
                break;
        case OP_LOADGLOBAL:
                top    = vm->top - 1;
                handle = top->value.i;
                assert(handle >= 0 && handle < MAX_NUM_HANDLES);

                struct obj_tensor_t* t = vm->handles[handle];
                if (t == NULL) {
                        return errNew(
                            "op OP_LOADGLOBAL: load a non-existed handle: %d",
                            handle);
                }

                printf("handle %d\n", handle);

                top->value.t = t;
                break;

        default:
                return errNew("unsupported opcode: %d", op);
        }

        return OK;
}
