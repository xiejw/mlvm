#ifndef MLVM_RANDOM_SPRNG_H_
#define MLVM_RANDOM_SPRNG_H_

#include "mlvm/random/sprng64.h"

typedef sprng64_t sprng_t;

#define sprng_create(seed) sprng64_create(seed)
#define sprng_free(prng) sprng64_free(prng)
#define sprng_split(prng) sprng64_split(prng)
#define sprng_next_int64(prng) sprng64_next_int64(prng)

#endif
