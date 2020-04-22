#ifndef MLVM_SPRNG_SPRNG_INTERFACE_H_
#define MLVM_SPRNG_SPRNG_INTERFACE_H_

#include "mlvm/sprng/sprng64.h"

/* Provides a public API for sprng. */
typedef sprng64_t sprng_t;

#define sprng_create(seed)     sprng64_create(seed)
#define sprng_free(prng)       sprng64_free(prng)
#define sprng_split(prng)      sprng64_split(prng)
#define sprng_next_int64(prng) sprng64_next_int64(prng)

#endif
