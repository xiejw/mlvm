#include "mlvm/random/normal.h"

#include <assert.h>
#include <float.h>
#include <math.h>
#include <stdlib.h>

/* unit of least precision */
static const double double_ulp_ = 1.0 / (1L << 53);
static const double two_pi_ = 2.0 * 3.141592653589793238;

void rng_normal(sprng_t* prng, size_t size, double* buffer) {
  size_t i;
  size_t num_seeds = size % 2 == 0 ? size : size + 1;
  double* uniforms = malloc(num_seeds * sizeof(double));

  assert(size > 0);

  for (i = 0; i < num_seeds;) {
    uint64_t seed = sprng_next_int64(prng);
    double u = (seed >> 11) * double_ulp_;
    if (i % 2 == 1 || u >= DBL_EPSILON) uniforms[i++] = u;
  }

  for (i = 0; i < size;) {
    double u_1 = uniforms[i];
    double u_2 = uniforms[i + 1];

    double theta = two_pi_ * u_1;
    double r = sqrt(-2.0 * log(u_2));

    buffer[i] = r * cos(theta);
    if (i + 1 < size) buffer[i + 1] = r * sin(theta);
    i += 2;
  }

  free(uniforms);
}
