#include "mlvm/testing/test.h"

#include "mlvm/sprng/sprng.h"

#define SIZE 4

static char* test_sprng_normal() {
  sprng_t* prng = sprng_create(456L);
  double   got[SIZE];
  double   expected[] = {0.447783, 0.324946, -0.638092, 2.241228};

  srng_standard_normal(prng, SIZE, got);

  ASSERT_ARRAY_CLOSE("Normal rv mismatch", expected, got, SIZE, 1e-6);

  sprng_free(prng);
  return NULL;
}

char* run_sprng_normal_test() {
  RUN_TEST(test_sprng_normal);
  return NULL;
}
