#include "mlvm/testing/test.h"

/* list_test.c */
extern char* run_sprng_normal_test();

char* run_sprng_test() {
  RUN_SUITE(run_sprng_normal_test);
  return NULL;
}
