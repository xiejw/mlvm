#include "mlvm/runtime/kernel/kernel.h"

#include <assert.h>

#include "mlvm/ir/tensor.h"

#define POS(x, stride_x, y, stride_y) ((x) * (stride_x) + (y) * (stride_y))

void kernel_matmul(tensor_t* output, tensor_t* arg_1, tensor_t* arg_2) {
  tensor_shape_t i, j, k, dim_i, dim_j, dim_k;
  tensor_size_t  stride_1_i, stride_1_j, stride_2_j, stride_2_k;
  /* We support trivial stride for output so far.*/
  assert(output->stride[1] == 1);

  assert(arg_1->rank == 2);
  assert(arg_1->rank == 2);
  assert(arg_1->shape[1] == arg_2->shape[0]);

  dim_i = arg_1->shape[0];
  dim_j = arg_1->shape[1];
  dim_k = arg_2->shape[1];

  stride_1_i = arg_1->stride[0];
  stride_1_j = arg_1->stride[1];
  stride_2_j = arg_2->stride[0];
  stride_2_k = arg_2->stride[1];

  for (i = 0; i < dim_i; i++) {
    for (k = 0; k < dim_k; k++) {
      double v = 0;
      for (j = 0; j < dim_j; j++) {
        v += arg_1->value[POS(i, stride_1_i, j, stride_1_j)] *
             arg_2->value[POS(j, stride_2_j, k, stride_2_k)];
      }
      /* TODO */
      output->value[POS(i, dim_k, k, 1)] = v;
    }
  }
}
