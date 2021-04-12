#include "op.h"

#include <assert.h>
#include <stdlib.h>

#include "rng/srng64_normal.h"

#define DEF_ELEWISE_OP(OP, op)                                                 \
        error_t vmOp##OP##F32(struct tensor_t* td, struct tensor_t* t1,        \
                              struct tensor_t* t2)                             \
        {                                                                      \
                assert(td->dtype == F32);                                      \
                assert(t1->dtype == F32);                                      \
                assert(t2->dtype == F32);                                      \
                                                                               \
                size_t size = t1->shape->size;                                 \
                if (size != td->shape->size) {                                 \
                        return errNew(                                         \
                            "Op " #OP                                          \
                            " support t1 and t2 same shape or t2 as scalar."); \
                }                                                              \
                                                                               \
                float32_t* o   = (float32_t*)td->data;                         \
                float32_t* lhs = (float32_t*)t1->data;                         \
                float32_t* rhs = (float32_t*)t2->data;                         \
                                                                               \
                if (t2->shape->size == size) {                                 \
                        for (size_t i = 0; i < size; i++) {                    \
                                o[i] = lhs[i] op rhs[i];                       \
                        }                                                      \
                } else if (t2->shape->size == 1) {                             \
                        float32_t s = rhs[0];                                  \
                        for (size_t i = 0; i < size; i++) {                    \
                                o[i] = lhs[i] op s;                            \
                        }                                                      \
                } else {                                                       \
                        return errNew(                                         \
                            "Op " #OP                                          \
                            " support t1 and t2 same shape or t2 as scalar."); \
                }                                                              \
                return OK;                                                     \
        }

DEF_ELEWISE_OP(Add, +)
DEF_ELEWISE_OP(Mul, *)
DEF_ELEWISE_OP(Minus, -)

#define DEF_ELEWISE_OP_S(OP, op)                                         \
        error_t vmOp##OP##SF32(struct tensor_t* td, struct tensor_t* t1, \
                               float32_t s)                              \
        {                                                                \
                assert(td->dtype == F32);                                \
                assert(t1->dtype == F32);                                \
                                                                         \
                size_t size = t1->shape->size;                           \
                assert(size == td->shape->size);                         \
                                                                         \
                float32_t* o   = (float32_t*)td->data;                   \
                float32_t* lhs = (float32_t*)t1->data;                   \
                                                                         \
                for (size_t i = 0; i < size; i++) {                      \
                        o[i] = lhs[i] op s;                              \
                }                                                        \
                return OK;                                               \
        }

DEF_ELEWISE_OP_S(Add, +)
DEF_ELEWISE_OP_S(Mul, *)
DEF_ELEWISE_OP_S(Minus, -)

error_t vmOpRngF32(struct tensor_t* dst, int mode,
                   const struct srng64_t* ori_rng)
{
        // make a copy to avoid advancing the ori_rng.
        struct srng64_t rng = *ori_rng;
        assert(mode == 0);
        assert(dst->dtype == F32);
        srng64StdNormalF(&rng, dst->shape->size, (float32_t*)dst->data);
        return OK;
}

error_t vmOpReduceF32(struct tensor_t* dst, struct tensor_t* t1, int mode)
{
        assert(mode == 0);  // sum
        assert(t1->dtype == F32);
        assert(dst->dtype == F32);
        assert(1 == dst->shape->size);

        float32_t  v    = 0;
        size_t     size = t1->shape->size;
        float32_t* data = (float32_t*)t1->data;
        for (size_t i = 0; i < size; i++) {
                v += data[i];
        }
        *(float32_t*)dst->data = v;
        return OK;
}
