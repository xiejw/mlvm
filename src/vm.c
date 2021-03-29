#include "vm.h"

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

error_t vmExec(enum opcode_t op, void* opt, int dst, int lhs, int rhs)
{
        return errNew("unimpl for vmExec");
}

void vmSync() {}
