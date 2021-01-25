#include "vm.h"

#include <stdio.h>
#include <string.h>

#include "opcode.h"

// -----------------------------------------------------------------------------
// internal prototypes.
// -----------------------------------------------------------------------------

#define STACK_INIT_SIZE 256
#define MAX_NUM_HANDLES 128

#define DEBUG_PRINT(x) printf(x)

typedef struct obj_tensor_item_t {
        struct obj_tensor_t*      item;
        struct obj_tensor_item_t* next;
} obj_tensor_item_t;

void* obj_tensor_pool = NULL;

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

vm_handle_t vmAllocTensor(struct vm_t* vm, int rank, int dims[])
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

        objTensorAllocAndCopy(t, /*buf=*/NULL);
        vm->handles[next_handle] = t;
        vm->size_used += size;
        return next_handle;
}

error_t vmDeallocTensor(struct vm_t* vm, vm_handle_t i)
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

// protocol
//
// - OP_HALT is the final one to halt vm.
// - Upon hatl, base must be reset to stack after execution.
//   - OP_RETURN could help.
// - all values after base must be tensors and will be copied as outputs.
// - all other tensors will be gc'ed.
error_t vmLaunch(struct vm_t* vm, vec_t(code_t) code,
                 vec_t(struct obj_tensor_t*) * outputs)
{
        error_t       err = OK;
        enum opcode_t op;
        code_t*       pc = code;

        while (1) {
                op = (enum opcode_t) * pc++;
                if (op != OP_HALT) {
                        if (handleOpCode(vm, &pc, op)) {
                                err = errEmitNote(
                                    "unexpected op handing error in vm.");
                                goto reset;
                        }
                } else {
                        DEBUG_PRINT("vm halt\n");
                        goto handle_outputs;
                }
        }

handle_outputs:

        // check invariance. we can loose this in future.
        if (vm->base != vm->stack) {
                err = errNew("illegal program. gabage left after execution.");
                goto reset;
        }

        // prepare the outputs;
        int count = vm->top - vm->base;
        assert(count >= 0);
        if (count == 0 || outputs == NULL) {
                goto reset;
        }

        // pass 1: check all left objs are tensors to avoid memory leak.
        for (struct obj_t* o = vm->base; o != vm->top; o++) {
                if (o->kind != OBJ_TENSOR) {
                        err = errNew(
                            "illegal program. return values must be all "
                            "tensors.");
                        goto reset;
                }
        }

        // pass 2: copy all tensors..
        int current_size = vecSize(*outputs);
        vecReserve(*outputs, current_size + count);
        vecSetSize(*outputs, current_size + count);

        for (int i = count - 1; i >= 0; i--) {
                int                  index = current_size + i;
                struct obj_t*        top   = --(vm->top);
                struct obj_tensor_t* t     = top->value.t;
                struct obj_tensor_t* dst   = objTensorNew(t->rank, t->dims);

                objTensorAllocAndCopy(dst, /*buf=*/t->buffer);
                (*outputs)[index] = dst;
        }

reset:
        vm->top  = vm->stack;
        vm->base = vm->stack;
        return err;
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
                top          = vm->top++;
                top->value.i = *((*pc)++);
                top->kind    = OBJ_INT;
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
                top->value.t = t;
                top->kind    = OBJ_TENSOR;
                break;

        default:
                return errNew("unsupported opcode: %d", op);
        }

        return OK;
}

// int objGC()
// {
//         if (obj_tensor_pool == NULL) return 0;
//
//         int                count = 0;
//         obj_tensor_item_t* p     = obj_tensor_pool;
//         obj_tensor_item_t* prev  = NULL;
//         while (p != NULL) {
//                 struct obj_tensor_t* item = p->item;
//                 if (item->mark) {
//                         item->mark = 0;
//                         prev       = p;
//                         p          = p->next;
//                 } else {
//                         if (prev == NULL) {
//                                 obj_tensor_pool = p->next;
//                         } else {
//                                 prev->next = p->next;
//                         }
//                         objTensorFree(item);
//
//                         obj_tensor_item_t* old_p;
//                         old_p = p;
//                         p     = p->next;
//                         free(old_p);
//                         count++;
//                 }
//         }
//         return count;
// }
//
