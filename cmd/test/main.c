#include <stdio.h>

#include "mlvm/testing/test.h"

#include "mlvm/lib/lib_test.h"
#include "mlvm/sprng/sprng_test.h"

typedef char* (*test_fn_t)();

typedef struct {
  char*     name;
  test_fn_t fn;
} test_suite_t;

test_suite_t test_suites[] = {
    {"mlvm/lib", run_lib_test},
    {"mlvm/sprng", run_sprng_test},
};

int main() {
  int i;
  int size = sizeof(test_suites) / sizeof(test_suite_t);
  for (i = 0; i < size; i++) {
    printf("Running suite: %s\n", test_suites[i].name);
    char* result = test_suites[i].fn();
    if (result != 0) {
      printf("\033[1;31mERROR: %s\033[0m\n", result); /* Red */
    } else {
      printf("\033[1;32mALL TESTS PASSED\033[0m\n"); /* Green */
    }
  }
  printf("Tests run: %d\n", tests_run);
  return 0;
}
