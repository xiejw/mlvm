#include "mlvm/runtime/kernel/kernel.h"

#include <assert.h>

#include "mlvm/ir/tensor.h"
#include "mlvm/runtime/kernel/macros.h"
#include "mlvm/runtime/kernel/shape_util.h"

void kernel_add(tensor_t* output, tensor_t* arg_1, tensor_t* arg_2) {
  tensor_size_t size = arg_1->size;

  assert(arg_1->size == arg_2->size);
  assert(kernel_shape_identical(arg_1, arg_2));  /* no broadcasting. */
  assert(kernel_stripe_identical(arg_1, arg_2)); /* naiive impl. */

  MLVM_KERNEL_ELEMENT_OP_PLAIN_LOOP(output, arg_1, arg_2, +, size);
}
