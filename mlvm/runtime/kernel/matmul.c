#include "mlvm/runtime/kernel/kernel.h"

#include <assert.h>

#include "mlvm/ir/tensor.h"

#define POS(x, y, dim_y) ((x) * (dim_y) + (y))

int kernel_matmul(tensor_t* output, tensor_t* arg_1, tensor_t* arg_2) {
  tensor_shape_t i, j, k, dim_i, dim_j, dim_k;
  /* We support trivial stride so far.*/
  assert(arg_1->stride[1] == 1);
  assert(arg_2->stride[1] == 1);

  assert(arg_1->rank == arg_2->rank);
  assert(arg_1->rank == 2);
  assert(arg_1->shape[1] == arg_2->shape[0]);

  dim_i = arg_1->shape[0];
  dim_j = arg_1->shape[1];
  dim_k = arg_2->shape[1];

  for (i = 0; i < dim_i; i++) {
    for (k = 0; k < dim_k; k++) {
      double v = 0;
      for (j = 0; j < dim_j; j++) {
        v += arg_1->value[POS(i, j, dim_j)] * arg_2->value[POS(j, k, dim_k)];
      }
      output->value[POS(i, k, dim_k)] = v;
    }
  }
  return 0;
}
