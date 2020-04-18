#include <stdio.h>

#include "mlvm/random/normal.h"
#include "mlvm/random/sprng.h"

int main() {
  sprng_t* prng = sprng_create(456L);
  double r_v[100];

  rng_normal(prng, 100, r_v);

  {
    int i;
    for (i = 0; i < 100; i++) {
      printf("%.3f ", r_v[i]);
      if (i % 10 == 9) printf("\n");
    }
  }

  sprng_free(prng);
  return 0;
}
