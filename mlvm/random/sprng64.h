#ifndef MLVM_RANDOM_SPRNG64_H_
#define MLVM_RANDOM_SPRNG64_H_

#include <stdint.h>

typedef struct {
  /* Internal fields. */
  uint64_t seed_;
  /* Coefficient for current level. */
  uint64_t gamma_;
  uint64_t next_split_;
} sprng64_t;

/*
 * seed should be 0 <= s <= gamma_prime_.
 * - `s` is the seed for next coefficient (gamma)
 */
extern sprng64_t* sprng64_create(uint64_t seed, uint64_t s);
extern void sprng64_free(sprng64_t* prng);
extern sprng64_t* sprng64_split(sprng64_t* prng);
extern uint64_t sprng64_next_int64(sprng64_t* prng);
extern uint32_t sprng64_next_int32(sprng64_t* prng);
/* 0 <= <=1 */
extern double sprng64_next_double(sprng64_t* prng);

#endif
