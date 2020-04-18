#ifndef MLVM_RANDOM_NORMAL_H_
#define MLVM_RANDOM_NORMAL_H_

#include <stddef.h>

#include "mlvm/random/sprng.h"

extern void rng_normal(sprng_t* prng, size_t size, double* buffer);

#endif
