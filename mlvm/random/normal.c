#include "mlvm/random/normal.h"

#include <assert.h>
#include <stdlib.h>

static const double double_ulp_ = 1.0 / (1L << 53);

void rng_normal(sprng_t* prng, size_t size, double* buffer) {
  uint64_t* seeds = malloc(size * sizeof(uint64_t));
  size_t i;

  assert(size > 0);
  for (i = 0; i < size; i++) {
    seeds[i] = sprng_next_int64(prng);
    buffer[i] = (seeds[i] >> 11) * double_ulp_;
  }

  free(seeds);
}
