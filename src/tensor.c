#include "vm.h"
#include "vm_internal.h"

#include <assert.h>
#include <stdlib.h>

// neg number if error. call site should clean error stack.
int
vmTensorNew(struct vm_t* vm, enum data_t dtype, struct shape_t* s)
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

error_t
vmTensorFree(struct vm_t* vm, int handle)
{
        vmReleaseHandle(vmGrabHandle(vm, handle));
        return OK;
}

// dtype and shape are optinoal (NULL).
error_t
vmTensorInfo(struct vm_t* vm, int handle, enum data_t* dtype,
             struct shape_t** shape)
{
        struct tensor_t* t = vmGrabHandle(vm, handle);

        assert(t->used);
        if (dtype != NULL) *dtype = t->dtype;
        if (shape != NULL) *shape = t->shape;
        return OK;
}

error_t
vmTensorData(struct vm_t* vm, int handle, void** data)
{
        struct tensor_t* t = vmGrabHandle(vm, handle);

        assert(t->used);
        *data = t->data;
        return OK;
}

void
vmTensorDump(sds_t* s, struct vm_t* vm, int handle)
{
        struct tensor_t* t  = vmGrabHandle(vm, handle);
        struct shape_t*  sp = t->shape;

        // shape
        sdsCatPrintf(s, "<");
        for (int i = 0; i < sp->rank; i++) {
                if (i != sp->rank - 1)
                        sdsCatPrintf(s, "%d, ", sp->dims[i]);
                else
                        sdsCatPrintf(s, "%d", sp->dims[i]);
        }
        sdsCatPrintf(s, "> ");

        // dtype
        switch (t->dtype) {
        case F32:
                sdsCatPrintf(s, "f32 ");
                break;
        case I32:
                sdsCatPrintf(s, "i32 ");
                break;
        default:
                sdsCatPrintf(s, " unknown_dtype ");
        }

        // data
#define print_data(dt, type_cast, fmt_str)                               \
        if (t->dtype == (dt)) {                                          \
                size_t size = sp->size;                                  \
                sdsCatPrintf(s, "[");                                    \
                for (size_t i = 0; i < size; i++) {                      \
                        if (i != size - 1)                               \
                                sdsCatPrintf(s, fmt_str ", ",            \
                                             ((type_cast)(t->data))[i]); \
                        else                                             \
                                sdsCatPrintf(s, fmt_str,                 \
                                             ((type_cast)(t->data))[i]); \
                        if (i >= 5 && i != size - 1) {                   \
                                sdsCatPrintf(s, "...");                  \
                                break;                                   \
                        }                                                \
                }                                                        \
                sdsCatPrintf(s, "]");                                    \
        }

        print_data(F32, float32_t*, "%.3f");
        print_data(I32, int32_t*, "%d");

#undef print_data
        return;
}
