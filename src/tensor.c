#include "vm.h"
#include "vm_internal.h"

#include <assert.h>
#include <stdlib.h>

typedef float float32_t;

// neg number if error. call site should clean error stack.
int vmNewT(struct vm_t* vm, enum data_t dtype, struct shape_t* s)
{
        void*            data;
        struct tensor_t* p = vm->handles;
        int              slot;
        for (slot = 0; slot < MAX_TENSOR_COUNT; slot++, p++) {
                if (p->used == 0) goto alloc;
        }
        return errNew("all handle spaces are used.");

alloc:
        switch (dtype) {
        case F32:
                data = malloc(s->size * sizeof(float32_t));
                break;
        case I32:
                data = malloc(s->size * sizeof(int32_t));
                break;
        default:
                return errNew("unsupported dtype for new tensor %d", dtype);
        }

        vmFillHandle(p, dtype, s, data);
        return slot;
}

error_t vmFreeT(struct vm_t* vm, int handle)
{
        assert(handle >= 0 && handle < MAX_TENSOR_COUNT);
        vmReleaseHandle(&vm->handles[handle]);
        return OK;
}

// dtype and shape are optinoal (NULL).
error_t vmFetchMetadata(struct vm_t* vm, int handle, enum data_t* dtype,
                        struct shape_t** shape)
{
        assert(handle >= 0 && handle < MAX_TENSOR_COUNT);
        struct tensor_t* t = &vm->handles[handle];

        assert(t->used);
        if (dtype != NULL) *dtype = t->dtype;
        if (shape != NULL) *shape = t->shape;
        return OK;
}

error_t vmFetchData(struct vm_t* vm, int handle, void** data)
{
        assert(handle >= 0 && handle < MAX_TENSOR_COUNT);
        struct tensor_t* t = &vm->handles[handle];

        assert(t->used);
        *data = t->data;
        return OK;
}
