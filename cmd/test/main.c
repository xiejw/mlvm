#include <stdio.h>

#include "mlvm/testing/test.h"

char *hello() {
  ASSERT_TRUE("should be true.", 1 == 0);
  return NULL;
}

char *run_suite() {
  RUN_TEST(hello);
  return NULL;
}

int main() {
  char *result = run_suite();
  if (result != 0) {
    printf("ERROR: %s\n", result);
  } else {
    printf("ALL TESTS PASSED\n");
  }
  printf("Tests run: %d\n", tests_run);
  return 0;
}
