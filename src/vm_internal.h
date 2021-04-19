#include <assert.h>
#include <stdlib.h>

// -----------------------------------------------------------------------------
// internal apis.
// -----------------------------------------------------------------------------

#define MAX_TENSOR_COUNT 128

struct list_t;

struct vm_t {
        // consider to use pages.
        struct tensor_t handles[MAX_TENSOR_COUNT];
        struct list_t*  shapes;
};

static inline struct tensor_t* vmGrabHandle(struct vm_t* vm, int handle)
{
        assert(handle >= 0 && handle < MAX_TENSOR_COUNT);
        return &vm->handles[handle];
}

// tensor_t is allocated in pages. so, we set the fields and mark as used.
static inline void vmFillHandle(struct tensor_t* t, enum data_t dtype,
                                struct shape_t* s, void* data)
{
        assert(!(t->used));
        assert(t->shape == NULL);
        assert(t->data == NULL);
        t->dtype = dtype;
        t->used  = 1;
        t->shape = spIncRef(s);
        t->data  = data;
}

// tensor_t is allocated in pages. so, we free the fields and mark as unused.
static inline void vmReleaseHandle(struct tensor_t* t)
{
        assert(t->used);
        free(t->data);
        spDecRef(t->shape);

        t->shape = NULL;
        t->data  = NULL;
        t->used  = 0;
}

// aux data structure.
struct list_t {
        void*          data;
        struct list_t* next;
};
