#include <stdio.h>

#include "mlvm/random/sprng.h"

int main() {
  sprng_t* prng = sprng_create(456L);
  printf("next double %lld\n", sprng_next_int64(prng));
  sprng_free(prng);
  return 0;
}
