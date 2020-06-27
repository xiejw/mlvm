#include "mlvm/testing/test.h"

/* normal_test.c */
extern char* run_sprng_normal_test();

char* run_sprng_suite() {
  RUN_SUITE(run_sprng_normal_test);
  return NULL;
}
