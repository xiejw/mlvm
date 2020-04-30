#include "mlvm/runtime/kernel/kernel.h"

#include <assert.h>

#include "mlvm/ir/tensor.h"

#define POS(x, stride_x, y, stride_y) ((x) * (stride_x) + (y) * (stride_y))

void kernel_matmul(tensor_t* output, tensor_t* arg_1, tensor_t* arg_2) {
  mlvm_uint_t i, j, k, dim_i, dim_j, dim_k;
  mlvm_size_t stride_1_i, stride_1_j;
  mlvm_size_t stride_2_j, stride_2_k;
  mlvm_size_t stride_output_i, stride_output_k;

  /* Assert all shapes. */
  assert(arg_1->rank == 2);
  assert(arg_1->rank == 2);
  assert(output->rank == 2);
  assert(arg_1->shape[1] == arg_2->shape[0]);
  assert(arg_1->shape[0] == output->shape[0]);
  assert(arg_2->shape[1] == output->shape[1]);

  /* Retrieve all dimensions. */
  dim_i = arg_1->shape[0];
  dim_j = arg_1->shape[1];
  dim_k = arg_2->shape[1];

  /* Retrieve all strides. */
  stride_1_i      = arg_1->stride[0];
  stride_1_j      = arg_1->stride[1];
  stride_2_j      = arg_2->stride[0];
  stride_2_k      = arg_2->stride[1];
  stride_output_i = output->stride[0];
  stride_output_k = output->stride[1];

  for (i = 0; i < dim_i; i++) {
    for (k = 0; k < dim_k; k++) {
      double v = 0;
      for (j = 0; j < dim_j; j++) {
        v += arg_1->value[POS(i, stride_1_i, j, stride_1_j)] *
             arg_2->value[POS(j, stride_2_j, k, stride_2_k)];
      }
      output->value[POS(i, stride_output_i, k, stride_output_k)] = v;
    }
  }
}
