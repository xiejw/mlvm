#include <stdio.h>
#include <unistd.h> /* STDOUT_FILENO */

#include "mlvm/ir/ir.h"
#include "mlvm/runtime/kernel/kernel.h"
#include "mlvm/sprng/sprng.h"

#define SIZE 100

int main() {
  sprng_t*       prng = sprng_create(456L);
  double         r_1[SIZE];
  uint32_t       shape[] = {10, 10};
  tensor_t*      t_1;
  ir_function_t* func;

  srng_standard_normal(prng, SIZE, r_1);

  t_1 = tensor_create(/*rank=*/2, shape, r_1, MLVM_ALIAS_VALUE);
  tensor_print(t_1, STDOUT_FILENO);

  func = ir_function_create("main");

  ir_function_free(func);
  tensor_free(t_1);

  sprng_free(prng);
  return 0;
}
