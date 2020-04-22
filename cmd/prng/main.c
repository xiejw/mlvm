#include <stdio.h>
#include <unistd.h>

#include "mlvm/ir/ir.h"
#include "mlvm/lib/list.h"
#include "mlvm/random/random.h"
#include "mlvm/runtime/kernel/kernel.h"

const int SIZE = 100;

int main() {
  sprng_t* prng = sprng_create(456L);
  double   r_1[SIZE];
  double   r_2[SIZE];
  double   r_3[SIZE];

  srng_standard_normal(prng, SIZE, r_1);
  srng_standard_normal(prng, SIZE, r_2);

  uint32_t  shape[] = {10, 10};
  tensor_t* t_1 =
      tensor_create(/*rank=*/2, /*shape=*/shape, r_1, MLVM_ALIAS_VALUE);
  tensor_print(t_1, STDOUT_FILENO);

  tensor_t* t_2 =
      tensor_create(/*rank=*/2, /*shape=*/shape, r_2, MLVM_ALIAS_VALUE);
  tensor_print(t_2, STDOUT_FILENO);

  tensor_t* t_3 =
      tensor_create(/*rank=*/2, /*shape=*/shape, r_3, MLVM_ALIAS_VALUE);

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

  {
    list_int_t lt;
    list_init(&lt);
    printf("List size %lld\n", list_size(&lt));
    list_append(&lt, 123);
    list_append(&lt, 456);
    printf("List 0 %d\n", list_get(&lt, 0));
    printf("List 1 %d\n", list_get(&lt, 1));
    printf("List size %lld\n", list_size(&lt));

    list_deinit(&lt);
  }
  return 0;
}
