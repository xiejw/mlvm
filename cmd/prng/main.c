#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

/*
 * rename prime
 * rename s in create.
 */

static const uint64_t gamma_prime_ = (1L << 56) - 5;
static const uint64_t gamma_gamma_ = 0x00281E2DBA6606F3L;
static const uint64_t default_seed_gamma_ = 0xBD24B73A95FB84D9L;
static const double double_ulp_ = 1.0 / (1L << 53);

#ifndef NDEBUG
#define debug_printf(fmt, ...) printf(fmt, __VA_ARGS__);
#else
#define debug_printf(fmt, ...)
#endif

typedef struct {
  /* Internal fields. */
  uint64_t seed_;
  /* Coefficient for current level. */
  uint64_t gamma_;
  uint64_t next_split_;
} sprng64_t;

static uint64_t sprng64_update(sprng64_t* prng, uint64_t s, uint64_t g) {
  /* Add g to s modulo George. */
  uint64_t p = s + g;
  return (p >= s) ? p : (p >= 0x800000000000000DL) ? p - 13L : (p - 13L) + g;
}

static uint64_t sprng64_mix64(uint64_t z) {
  z = ((z ^ (z >> 33)) * 0xff51afd7ed558ccdL);
  z = ((z ^ (z >> 33)) * 0xc4ceb9fe1a85ec53L);
  return z ^ (z >> 33);
}

static uint64_t sprng64_mix56(uint64_t z) {
  z = ((z ^ (z >> 33)) * 0xff51afd7ed558ccdL) & 0x00FFFFFFFFFFFFFFL;
  z = ((z ^ (z >> 33)) * 0xc4ceb9fe1a85ec53L) & 0x00FFFFFFFFFFFFFFL;
  return z ^ (z >> 33);
}

static uint64_t sprng64_next_raw64(sprng64_t* prng) {
  /* Advanced one more coefficient at current level. */
  return (prng->seed_ = sprng64_update(prng, prng->seed_, prng->gamma_));
}

/*
 * seed should be 0 <= s <= gamma_prime_.
 * - `s` is the seed for next coefficient (gamma)
 */
sprng64_t* sprng64_create(uint64_t seed, uint64_t s) {
  sprng64_t* prng = malloc(sizeof(sprng64_t));
  prng->seed_ = seed;
  s += gamma_gamma_;
  if (s >= gamma_prime_) s -= gamma_prime_;
  prng->gamma_ = sprng64_mix56(s) + 13;
  prng->next_split_ = s; /* used for next gamme coefficient. */
  return prng;
}

sprng64_t* sprng64_split(sprng64_t* prng) {
  uint64_t seed = sprng64_next_raw64(prng);
  uint64_t s = prng->next_split_;
  return sprng64_create(seed, s);
}

void sprng64_free(sprng64_t* prng) { free(prng); }

int main() {
  debug_printf("prime_ %llx\n", gamma_prime_);
  debug_printf("gamma_ %llx\n", gamma_gamma_);
  debug_printf("default seed gamma_ %llx\n", default_seed_gamma_);
  debug_printf("double ulp_ %.54f\n", double_ulp_);
  return 0;
}
