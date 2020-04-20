#include <stdio.h>

#include "mlvm/random/normal.h"
#include "mlvm/random/sprng.h"

const int SIZE = 100;

int main() {
  sprng_t* prng = sprng_create(456L);
  double   r_v[SIZE];

  rng_standard_normal(prng, SIZE, r_v);

  {
    int i;
    for (i = 0; i < SIZE; i++) {
      printf("%6.3f  ", r_v[i]);
      if (i % 10 == 9) printf("\n");
    }
  }

  sprng_free(prng);
  return 0;
}
