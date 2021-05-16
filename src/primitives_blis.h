// Provides inline BLIS kernels for BLAS.
#include "blis.h"

static inline void
vmBlisMatmul(int m, int n, int k, float* a, float* b, float* c)
{
        float32_t zero = 0;
        float32_t one  = 1;
        bli_sgemm(
            /*trans_a=*/BLIS_NO_TRANSPOSE,
            /*trans_b=*/BLIS_NO_TRANSPOSE,
            /*m=*/m,
            /*n=*/n,
            /*k=*/k,
            /*alpha=*/&one, a, k, 1, b, n, 1, /*beta=*/&zero, c, n, 1);
}

static inline void
vmBlisMatmulTR(int m, int n, int k, float* a, float* b, float* c)
{
        float32_t zero = 0;
        float32_t one  = 1;
        bli_sgemm(
            /*trans_a=*/BLIS_NO_TRANSPOSE,
            /*trans_b=*/BLIS_TRANSPOSE,
            /*m=*/m,
            /*n=*/n,
            /*k=*/k,
            /*alpha=*/&one, a, k, 1, b, k, 1, /*beta=*/&zero, c, n, 1);
}

static inline void
vmBlisMatmulTL(int m, int n, int k, float* a, float* b, float* c)
{
        float32_t zero = 0;
        float32_t one  = 1;
        bli_sgemm(
            /*trans_a=*/BLIS_TRANSPOSE,
            /*trans_b=*/BLIS_NO_TRANSPOSE,
            /*m=*/m,
            /*n=*/n,
            /*k=*/k,
            /*alpha=*/&one, a, m, 1, b, n, 1, /*beta=*/&zero, c, n, 1);
}

error_t
vmOpMinusF32(struct tensor_t* td, struct tensor_t* t1, struct tensor_t* t2)
{
        assert(td->dtype == F32);
        assert(t1->dtype == F32);
        assert(t2->dtype == F32);

        size_t size = t1->shape->size;
        if (size != td->shape->size) {
                return errNew(
                    "Op MINUS support t1 and t2 same shape or t2 as scalar.");
        }

        float32_t* o   = (float32_t*)td->data;
        float32_t* lhs = (float32_t*)t1->data;
        float32_t* rhs = (float32_t*)t2->data;

        size_t t2_size = t2->shape->size;

        float32_t zeros      = 0;
        float32_t minus_ones = -1;
        if (t2_size == size || t2_size == 1) {
                size_t stride = t2_size == 1 ? 0 : 1;
                if (t1 == td) {
                        // special case for self minus
                        // y = y - x
                        bli_ssubv(
                            /*conjx=*/BLIS_NO_CONJUGATE, size, /*x=*/rhs,
                            stride,
                            /*y=*/o, 1);
                } else {
                        // z = y + alphax * x + alphay * y.
                        bli_saxpy2v(BLIS_NO_CONJUGATE, BLIS_NO_CONJUGATE, size,
                                    /*alphax=*/&minus_ones, /*alphay=*/&zeros,
                                    /*x=*/rhs, stride, /*y=*/lhs, 1, o, 1);
                }
                return OK;
        }

        if (size % t2_size == 0) {
                size_t loop_c = size / t2_size;
                size_t offset = 0;
                for (size_t c = 0; c < loop_c; c++) {
                        // z = y + alphax * x + alphay * y.
                        bli_saxpy2v(
                            BLIS_NO_CONJUGATE, BLIS_NO_CONJUGATE, t2_size,
                            /*alphax=*/&minus_ones, /*alphay=*/&zeros,
                            /*x=*/rhs, 1, /*y=*/lhs + offset, 1, o + offset, 1);

                        offset += t2_size;
                }
                return OK;
        }
        return errNew("Op_MINUS support t1s==t2s, t1s%%t2s==0 or t2s==1.");
}
