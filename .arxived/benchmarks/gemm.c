#include <stdio.h>

#include "blis.h"

#define K 4096
#define N 4096

int gemm1(float *restrict C, float *restrict A, float *restrict B, int start,
          int end) {
  int i, j, k;

  for (i = start; i < end; i++) {
    for (k = 0; k < K; k++) {
      float a      = A[i * K + k];
      int   indexC = i * N;
      int   indexB = k * N;

      for (j = 0; j < N; j++) {
        C[indexC++] += a * B[indexB++];
      }
    }
  }
  return 0;
}

int gemm2(float *restrict C, float *restrict A, float *restrict B, int start,
          int end) {
  float alpha = 1.0;
  float beta  = 1.0;
  int   m     = end - start;

  float *a = &A[start * K];
  float *c = &C[start * N];

  bli_sgemm(BLIS_NO_TRANSPOSE, BLIS_NO_TRANSPOSE, m, N, K, &alpha, a, K, 1, B,
            N, 1, &beta, c, N, 1);
  return 0;
}
