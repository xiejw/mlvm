#include <stdio.h>

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
