#include "op.h"

#include <assert.h>
#include <stdlib.h>

#include "rng/srng64_normal.h"

error_t vmOpAddF32(struct tensor_t* td, struct tensor_t* t1,
                   struct tensor_t* t2)
{
        assert(td->dtype == F32);
        assert(t1->dtype == F32);
        assert(t2->dtype == F32);

        size_t size = t1->shape->size;
        assert(size == t2->shape->size);
        assert(size == td->shape->size);

        float32_t* o   = (float32_t*)td->data;
        float32_t* lhs = (float32_t*)t1->data;
        float32_t* rhs = (float32_t*)t2->data;

        for (size_t i = 0; i < size; i++) {
                o[i] = lhs[i] + rhs[i];
        }
        return OK;
}

error_t vmOpcRngF32(struct tensor_t* dst, int mode,
                    const struct srng64_t* ori_rng)
{
        assert(mode == 0);
        assert(dst->dtype == F32);
        // make a copy to avoid advancing the ori_rng.
        struct srng64_t rng = *ori_rng;
        srng64StdNormalF(&rng, dst->shape->size, (float32_t*)dst->data);
        return OK;
}
