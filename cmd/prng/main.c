#include <stdio.h>
#include <unistd.h>

#include "mlvm/ir/tensor.h"
#include "mlvm/random/normal.h"
#include "mlvm/random/sprng.h"

const int SIZE = 100;

int main() {
  sprng_t* prng = sprng_create(456L);
  double   r_v[SIZE];

  rng_standard_normal(prng, SIZE, r_v);

  uint32_t shape[] = {10, 10};

  mlvm_tensor_t* tensor =
      mlvm_tensor_create(/*rank=*/2, /*shape=*/shape, r_v, MLVM_ALIAS_VALUE);

  mlvm_tensor_print(tensor, STDOUT_FILENO);

  mlvm_tensor_free(tensor);

  sprng_free(prng);
  return 0;
}
