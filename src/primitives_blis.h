// Provides inline BLIS kernels for BLAS.
#include "blis.h"

static inline void
vmBlisMatmulF32(int m, int n, int k, float* a, float* b, float* c)
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
vmBlisMatmulTRF32(int m, int n, int k, float* a, float* b, float* c)
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
vmBlisMatmulTLF32(int m, int n, int k, float* a, float* b, float* c)
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

static inline void
vmBlisMinusF32(float32_t* o, float32_t* lhs, float32_t* rhs, size_t size,
               size_t rhs_size)
{
        float32_t zeros      = 0;
        float32_t minus_ones = -1;
        if (rhs_size == size || rhs_size == 1) {
                size_t stride = rhs_size == 1 ? 0 : 1;
                if (o == lhs) {
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
                return;
        }

        if (size % rhs_size == 0) {
                size_t loop_c = size / rhs_size;
                size_t offset = 0;
                for (size_t c = 0; c < loop_c; c++) {
                        // z = y + alphax * x + alphay * y.
                        bli_saxpy2v(
                            BLIS_NO_CONJUGATE, BLIS_NO_CONJUGATE, rhs_size,
                            /*alphax=*/&minus_ones, /*alphay=*/&zeros,
                            /*x=*/rhs, 1, /*y=*/lhs + offset, 1, o + offset, 1);

                        offset += rhs_size;
                }
                return;
        }

        errFatalAndExit(
            "invalid size for vmBlisMinusF32. size: %d, rhs_size: %d\n", size,
            rhs_size);
}

static inline void
vmBlisMulSF32(float32_t* o, float32_t* lhs, float32_t v, size_t size)
{
        // y = alpha * x
        bli_sscal2v(BLIS_NO_CONJUGATE, size, /*alpha=*/&v,
                    /*x=*/lhs, 1, /*y=*/o, 1);
}
