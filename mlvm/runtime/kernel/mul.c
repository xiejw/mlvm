#include "mlvm/runtime/kernel/kernel.h"

#include <assert.h>

#include "mlvm/ir/tensor.h"
#include "mlvm/runtime/kernel/shape_util.h"

int kernel_mul(tensor_t* output, tensor_t* arg_1, tensor_t* arg_2) {
  tensor_size_t i;
  tensor_size_t size = arg_1->size;

  assert(arg_1->size == arg_2->size);
  assert(kernel_shape_identical(arg_1, arg_2));  /* no broadcasting. */
  assert(kernel_stripe_identical(arg_1, arg_2)); /* naiive impl. */

  for (i = 0; i < size; i++) {
    output->value[i] = arg_1->value[i] * arg_2->value[i];
  }
  return 0;
}
