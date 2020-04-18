#ifndef MLVM_RANDOM_SPRNG64_H_
#define MLVM_RANDOM_SPRNG64_H_

#include <stdint.h>

typedef struct {
  /* Internal fields. */
  uint64_t seed_;
  uint64_t gamma_;
  uint64_t next_gamma_seed_;
} sprng64_t;

extern sprng64_t* sprng64_create(uint64_t seed, uint64_t s);
extern void sprng64_free(sprng64_t* prng);
extern sprng64_t* sprng64_split(sprng64_t* prng);

extern uint64_t sprng64_next_int64(sprng64_t* prng);
extern uint32_t sprng64_next_int32(sprng64_t* prng);
extern double sprng64_next_double(sprng64_t* prng);

#endif
