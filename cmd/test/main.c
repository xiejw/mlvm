#include <stdio.h>

#include "mlvm/testing/test.h"

#include "mlvm/lib/lib_test.h"
#include "mlvm/random/random_test.h"

typedef char* (*test_fn_t)();

typedef struct {
  char*     name;
  test_fn_t fn;
} test_suite_t;

test_suite_t test_suites[] = {
    {"mlvm/lib", run_lib_test},
    {"mlvm/random", run_sprng_test},
};

int main() {
  int i;
  int size = sizeof(test_suites) / sizeof(test_suite_t);
  for (i = 0; i < size; i++) {
    printf("Running suite: %s\n", test_suites[i].name);
    char* result = test_suites[i].fn();
    if (result != 0) {
      printf("ERROR: %s\n", result);
    } else {
      printf("ALL TESTS PASSED\n");
    }
  }
  printf("Tests run: %d\n", tests_run);
  return 0;
}
