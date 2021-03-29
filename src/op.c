#include "op.h"

#include "vm_internal.h"

#include <assert.h>
#include <stdlib.h>

error_t opAdd(struct tensor_t* dst, struct tensor_t* t1, struct tensor_t* t2)
{
        size_t size = t1->shape->size;
        assert(size == t2->shape->size);
        assert(size == dst->shape->size);

        assert(dst->dtype == F32);
        assert(t1->dtype == F32);
        assert(t2->dtype == F32);

        float32_t* o   = (float32_t*)dst->data;
        float32_t* lhs = (float32_t*)t1->data;
        float32_t* rhs = (float32_t*)t2->data;

        for (size_t i = 0; i < size; i++) {
                o[i] = lhs[i] + rhs[i];
        }
        return OK;
}
