#include <stdio.h>

#include "mlvm/random/normal.h"
#include "mlvm/random/sprng.h"

const int SIZE = 100;

int main() {
  sprng_t* prng = sprng_create(456L);
  double r_v[SIZE];

  rng_normal(prng, SIZE, r_v);

  {
    int i;
    for (i = 0; i < SIZE; i++) {
      printf("%.3f\n", r_v[i]);
    }
  }

  sprng_free(prng);
  return 0;
}
