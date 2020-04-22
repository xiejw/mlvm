#ifndef MLVM_RANDOM_NORMAL_H_
#define MLVM_RANDOM_NORMAL_H_

#include <stddef.h>

#include "mlvm/sprng/sprng_api.h"

extern void srng_standard_normal(sprng_t* prng, size_t size, double* buffer);

#endif
