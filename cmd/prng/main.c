#include <stdint.h>
#include <stdio.h>

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
  uint64_t gamma_;
  uint64_t next_split_;
} sprng64_t;

static uint64_t sprng64_update(sprng64_t* prng, uint64_t s, uint64_t g) {
  /* Add g to s modulo George. */
  uint64_t p = s + g;
  return (p >= s) ? p : (p >= 0x800000000000000DL) ? p - 13L : (p - 13L) + g;
}

static uint64_t sprng64_mix56(sprng64_t* prng, uint64_t z) {
  z = ((z ^ (z >> 33)) * 0xff51afd7ed558ccdL) & 0x00FFFFFFFFFFFFFFL;
  z = ((z ^ (z >> 33)) * 0xc4ceb9fe1a85ec53L) & 0x00FFFFFFFFFFFFFFL;
  return z ^ (z >> 33);
}

static uint64_t sprng64_next_raw64(sprng64_t* prng) {
  return (prng->seed_ = sprng64_update(prng, prng->seed_, prng->gamma_));
}

int main() {
  debug_printf("prime_ %llx\n", gamma_prime_);
  debug_printf("gamma_ %llx\n", gamma_gamma_);
  debug_printf("default seed gamma_ %llx\n", default_seed_gamma_);
  debug_printf("double ulp_ %.54f\n", double_ulp_);
  return 0;
}
