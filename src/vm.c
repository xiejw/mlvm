#include "vm.h"

#include "op.h"
#include "vm_internal.h"

#include <string.h>  // memset

struct vm_t* vmNew()
{
        size_t       size = sizeof(struct vm_t);
        struct vm_t* vm   = malloc(size);
        memset(vm, 0, size);
        return vm;
}

void vmFree(struct vm_t* vm)
{
        for (int i = 0; i < MAX_TENSOR_COUNT; i++) {
                struct tensor_t* t = &vm->handles[i];
                if (t->used) vmReleaseHandle(t);
        }
        free(vm);
}

error_t vmExec(struct vm_t* vm, enum opcode_t op, void* opt, int dst, int lhs,
               int rhs)
{
        struct tensor_t* td = vmGrabHandle(vm, dst);
        struct tensor_t* t1 = vmGrabHandle(vm, lhs);
        struct tensor_t* t2 = vmGrabHandle(vm, rhs);

        switch (op) {
        case OP_ADD:
                assert(opt == NULL);
                return opAdd(td, t1, t2);

        default:
                return errNewWithNote(ENOTIMPL, "unimpl for vmExec");
        }
}

void vmSync() {}
