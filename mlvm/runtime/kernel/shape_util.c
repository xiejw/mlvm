#include "mlvm/runtime/kernel/shape_util.h"

#include <assert.h>

extern int kernel_stripe_identical(tensor_t* arg_1, tensor_t* arg_2) {
  tensor_shape_t i;
  tensor_shape_t rank     = arg_1->rank;
  tensor_size_t *stride_1 = arg_1->stride, *stride_2 = arg_2->stride;

  assert(arg_1->rank == arg_2->rank);
  for (i = 0; i < rank; i++) {
    if (stride_1[i] != stride_2[i]) return 0;
  }
  return 1;
}

extern int kernel_shape_identical(tensor_t* arg_1, tensor_t* arg_2) {
  tensor_shape_t  i;
  tensor_shape_t  rank    = arg_1->rank;
  tensor_shape_t *shape_1 = arg_1->shape, *shape_2 = arg_2->shape;

  assert(arg_1->rank == arg_2->rank);
  for (i = 0; i < rank; i++) {
    if (shape_1[i] != shape_2[i]) return 0;
  }
  return 1;
}
