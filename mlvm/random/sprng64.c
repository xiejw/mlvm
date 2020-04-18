#include "mlvm/random/sprng64.h"

#include <stdlib.h>

/*
 * rename prime
 * rename s in create. split?
 * why 56
 * why 53
 */

static const uint64_t gamma_prime_ = (1L << 56) - 5;
/*
 * The coefficient to generate prng for coefficient.
 *
 * By using prng to generate coefficient for different, a fixed table is
 * avoided. The algorithrm is based on DOTMIX with a length-1 pedigee.
 */
static const uint64_t gamma_gamma_ = 0x00281E2DBA6606F3L;
static const double double_ulp_ = 1.0 / (1L << 53);

static uint64_t sprng64_update(uint64_t s, uint64_t g) {
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
  return (prng->seed_ = sprng64_update(prng->seed_, prng->gamma_));
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
  /* Advance the seed for current `prng` and use it for the seed for splitted
   * branch. */
  uint64_t seed = sprng64_next_raw64(prng);
  uint64_t s = prng->next_split_;
  return sprng64_create(seed, s);
}

void sprng64_free(sprng64_t* prng) { free(prng); }

uint64_t sprng64_next_int64(sprng64_t* prng) {
  return sprng64_mix64(sprng64_next_raw64(prng));
}

uint32_t sprng64_next_int32(sprng64_t* prng) {
  return (uint32_t)(sprng64_next_int64(prng));
}

/* 0 <= <=1 */
double sprng64_next_double(sprng64_t* prng) {
  return (sprng64_next_int64(prng) >> 11) * double_ulp_;
}
