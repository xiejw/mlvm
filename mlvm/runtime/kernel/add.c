#include "mlvm/runtime/kernel/kernel.h"

#include <assert.h>

#include "mlvm/ir/tensor.h"

int kernel_add(tensor_t* output, tensor_t* arg_1, tensor_t* arg_2) {
  uint64_t i;
  uint64_t size = arg_1->size;

  assert(arg_1->size == arg_2->size);
  for (i = 0; i < size; i++) {
    output->value[i] = arg_1->value[i] + arg_2->value[i];
  }
  return 0;
}
