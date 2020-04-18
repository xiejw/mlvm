#include <stdio.h>

#include "mlvm/random/sprng64.h"

int main() {
  sprng64_t* prng = sprng64_create(456L);
  printf("next double %.54f\n", sprng64_next_double(prng));
  sprng64_free(prng);

  return 0;
}
