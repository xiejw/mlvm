#include <stdio.h>

#include "mlvm/ir/ir.h"
#include "mlvm/runtime/kernel/kernel.h"
#include "mlvm/sprng/sprng.h"

#define SIZE 100
#define SHAPE \
  { 10, 10 }
#define RANK 2

tensor_t* create_a_random_tensor(sprng_t* prng) {
  double*   value;
  uint32_t  shape[] = SHAPE;
  tensor_t* tensor;
  value = malloc(SIZE * sizeof(double)); /* This value moved into the tensor. */
  srng_standard_normal(prng, SIZE, value);
  tensor = tensor_create(RANK, shape, value, MLVM_MOVE_VALUE);
  return tensor;
}

int build_func(ir_function_t* func, sprng_t* prng) {
  tensor_t* tensor;
  tensor = create_a_random_tensor(prng);
  if (NULL == ir_function_append_constant(func, tensor, MLVM_MOVE_VALUE)) {
    return -1;
  }
  tensor_free(tensor);

  ir_function_print(func, 1);
  return 0;
}

int main() {
  sprng_t* prng = sprng_create(456L);

  ir_function_t* func;
  func = ir_function_create("main");

  if (build_func(func, prng)) fprintf(stderr, "Unexpected error.");

  ir_function_free(func);
  sprng_free(prng);
  return 0;
}
