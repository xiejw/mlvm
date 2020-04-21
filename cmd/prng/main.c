#include <stdio.h>
#include <unistd.h>

#include "mlvm/ir/tensor.h"
#include "mlvm/random/normal.h"
#include "mlvm/random/sprng.h"
#include "mlvm/runtime/kernel.h"

const int SIZE = 100;

int main() {
  sprng_t* prng = sprng_create(456L);
  double   r_1[SIZE];
  double   r_2[SIZE];
  double   r_3[SIZE];

  rng_standard_normal(prng, SIZE, r_1);
  rng_standard_normal(prng, SIZE, r_2);

  uint32_t       shape[] = {10, 10};
  mlvm_tensor_t* t_1 =
      mlvm_tensor_create(/*rank=*/2, /*shape=*/shape, r_1, MLVM_ALIAS_VALUE);
  mlvm_tensor_print(t_1, STDOUT_FILENO);

  mlvm_tensor_t* t_2 =
      mlvm_tensor_create(/*rank=*/2, /*shape=*/shape, r_2, MLVM_ALIAS_VALUE);
  mlvm_tensor_print(t_2, STDOUT_FILENO);

  mlvm_tensor_t* t_3 =
      mlvm_tensor_create(/*rank=*/2, /*shape=*/shape, r_3, MLVM_ALIAS_VALUE);

  mlvm_runtime_kernel_add(t_3, t_1, t_2);
  mlvm_tensor_print(t_3, STDOUT_FILENO);

  mlvm_tensor_free(t_1);
  mlvm_tensor_free(t_2);
  mlvm_tensor_free(t_3);

  sprng_free(prng);
  return 0;
}
