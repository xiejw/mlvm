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
  uint64_t seed_;
} splittable64_t;

int main() {
  debug_printf("prime_ %llx\n", gamma_prime_);
  debug_printf("gamma_ %llx\n", gamma_gamma_);
  debug_printf("default seed gamma_ %llx\n", default_seed_gamma_);
  debug_printf("double ulp_ %.54f\n", double_ulp_);
  return 0;
}
