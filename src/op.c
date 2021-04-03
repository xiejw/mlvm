#include "op.h"

#include <assert.h>
#include <stdlib.h>

struct opopt_t* vmOptNew()
{
        struct opopt_t* opt = malloc(sizeof(struct opopt_t));
        opt->ref_count      = 0;
        opt->mode           = 0;
        return opt;
}

struct opopt_t* vmOptIncRef(struct opopt_t* opt)
{
        assert(opt->ref_count > 0);
        opt->ref_count++;
        return opt;
}

struct opopt_t* vmOptDecRef(struct opopt_t* opt)
{
        if (--(opt->ref_count) == 0) {
                free(opt);
                return NULL;
        }
        return opt;
}

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
