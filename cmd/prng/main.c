#include <stdio.h>
#include <unistd.h> /* STDOUT_FILENO */

#include "mlvm/ir/ir.h"
#include "mlvm/runtime/kernel/kernel.h"
#include "mlvm/sprng/sprng.h"

#define SIZE 100

int main() {
  sprng_t*  prng = sprng_create(456L);
  double    r_1[SIZE];
  double    r_2[SIZE];
  double    r_3[SIZE];
  uint32_t  shape[] = {10, 10};
  tensor_t *t_1, *t_2, *t_3;

  srng_standard_normal(prng, SIZE, r_1);
  srng_standard_normal(prng, SIZE, r_2);

  t_1 = tensor_create(/*rank=*/2, /*shape=*/shape, r_1, MLVM_ALIAS_VALUE);
  tensor_print(t_1, STDOUT_FILENO);

  t_2 = tensor_create(/*rank=*/2, /*shape=*/shape, r_2, MLVM_ALIAS_VALUE);
  tensor_print(t_2, STDOUT_FILENO);

  t_3 = tensor_create(/*rank=*/2, /*shape=*/shape, r_3, MLVM_ALIAS_VALUE);

  kernel_add(t_3, t_1, t_2);
  tensor_print(t_3, STDOUT_FILENO);

  kernel_mul(t_3, t_1, t_2);
  tensor_print(t_3, STDOUT_FILENO);

  kernel_matmul(t_3, t_1, t_2);
  tensor_print(t_3, STDOUT_FILENO);

  tensor_free(t_1);
  tensor_free(t_2);
  tensor_free(t_3);

  sprng_free(prng);
  return 0;
}
