#include "primitives.h"

#include <assert.h>
#include <math.h>
#include <stdlib.h>

#include "rng/srng64_normal.h"

#define PLUS(x, y) ((x) + (y))
#define MULT(x, y) ((x) * (y))
#define MINU(x, y) ((x) - (y))
#define MAXI(x, y) ((x) > (y) ? (x) : (y))
#define CMPL(x, y) ((x) > (y) ? (1) : (0))

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
                size_t t2_size = t2->shape->size;                              \
                                                                               \
                if (t2_size == size) {                                         \
                        for (size_t i = 0; i < size; i++) {                    \
                                o[i] = op(lhs[i], rhs[i]);                     \
                        }                                                      \
                } else if (t2_size == 1) {                                     \
                        float32_t s = rhs[0];                                  \
                        for (size_t i = 0; i < size; i++) {                    \
                                o[i] = op(lhs[i], s);                          \
                        }                                                      \
                } else if (size % t2_size == 0) {                              \
                        size_t loop_c = size / t2_size;                        \
                        for (size_t c = 0; c < loop_c; c++) {                  \
                                size_t offset = c * t2_size;                   \
                                for (size_t i = 0; i < t2_size; i++) {         \
                                        o[offset + i] =                        \
                                            op(lhs[offset + i], rhs[i]);       \
                                }                                              \
                        }                                                      \
                } else {                                                       \
                        return errNew(                                         \
                            "Op " #OP                                          \
                            " support t1s==t2s, t1s%%t2s==0 or t2s==1.");      \
                }                                                              \
                return OK;                                                     \
        }

DEF_ELEWISE_OP(Add, PLUS)
DEF_ELEWISE_OP(Mul, MULT)
DEF_ELEWISE_OP(Minus, MINU)
DEF_ELEWISE_OP(Max, MAXI)
DEF_ELEWISE_OP(CmpL, CMPL)

#undef DEF_ELEWISE_OP

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
                        o[i] = op(lhs[i], s);                            \
                }                                                        \
                return OK;                                               \
        }

DEF_ELEWISE_OP_S(Add, PLUS)
DEF_ELEWISE_OP_S(Mul, MULT)
DEF_ELEWISE_OP_S(Minus, MINU)
DEF_ELEWISE_OP_S(Max, MAXI)
DEF_ELEWISE_OP_S(CmpL, CMPL)

#undef DEF_ELEWISE_OP_S

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

error_t vmOpMatmulF32(struct tensor_t* td, struct tensor_t* t1,
                      struct tensor_t* t2)
{
        assert(td != t1);
        assert(td != t2);
        assert(td->dtype == F32);
        assert(t1->dtype == F32);
        assert(t2->dtype == F32);
        assert(td->shape->rank == 2);
        assert(t1->shape->rank == 2);
        assert(t2->shape->rank == 2);

        // pm,mq -> pq
        int p = td->shape->dims[0];
        int q = td->shape->dims[1];
        int m = t1->shape->dims[1];

        if (p != t1->shape->dims[0] || m != t2->shape->dims[0] ||
            q != t2->shape->dims[1]) {
                return errNew("invalid matmul shape: %d/%d,%d/%d->%d/%d.",
                              t1->shape->dims[0], t1->shape->dims[1],
                              t2->shape->dims[0], t2->shape->dims[1],
                              td->shape->dims[0], td->shape->dims[1]);
        }

        float32_t* o   = (float32_t*)td->data;
        float32_t* lhs = (float32_t*)t1->data;
        float32_t* rhs = (float32_t*)t2->data;

        // stupid impl.
        for (int i = 0; i < p; i++) {
                for (int j = 0; j < q; j++) {
                        float32_t v = 0;
                        for (int k = 0; k < m; k++) {
                                v += lhs[i * m + k] * rhs[k * q + j];
                        }
                        o[i * q + j] = v;
                }
        }

        return OK;
}

error_t vmOpLossSCELF32(struct tensor_t* td, struct tensor_t* t1,
                        struct tensor_t* t2)
{
        assert(td != t1);
        assert(td != t2);
        assert(td->dtype == F32);
        assert(t1->dtype == F32);
        assert(t2->dtype == F32);
        assert(td->shape->rank == 1);
        assert(t1->shape->rank == 2);
        assert(t2->shape->rank == 2);
        int bs = td->shape->dims[0];
        int ft = t1->shape->dims[1];

        if (t1->shape->dims[0] != bs || t2->shape->dims[0] != bs ||
            t2->shape->dims[1] != ft) {
                return errNew("invalid LossSCEL shape: %d/%d,%d/%d->%d.",
                              t1->shape->dims[0], t1->shape->dims[1],
                              t2->shape->dims[0], t2->shape->dims[1],
                              td->shape->dims[0]);
        }
        float32_t* loss = (float32_t*)td->data;
        float32_t* y    = (float32_t*)t1->data;
        float32_t* o    = (float32_t*)t2->data;

        for (size_t i = 0; i < bs; i++) {
                // find max exp and shift the value to become stable.
                size_t  offset  = i * ft;
                float_t max_o_k = o[offset];
                for (size_t k = 1; k < ft; k++) {
                        float32_t o_k = o[offset + k];
                        if (o_k > max_o_k) max_o_k = o_k;
                }

                // real formular
                float32_t sum = 0.0;
                float32_t l   = 0.0;
                for (size_t k = 0; k < ft; k++) {
                        float32_t o_k = o[offset + k] - max_o_k;
                        sum += exp(o_k);
                        l -= y[offset + k] * o_k;
                }

                loss[i] = l + log(sum);
        }
        return OK;
}
