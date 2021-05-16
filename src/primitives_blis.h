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
