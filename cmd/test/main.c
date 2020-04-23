#include <stdio.h>

#include "mlvm/testing/test.h"

#include "mlvm/lib/test_suite.h"
#include "mlvm/sprng/test_suite.h"

typedef char* (*test_fn_t)();

typedef struct {
  char*     name;
  test_fn_t fn;
} test_suite_t;

/*
 * Testing Hierarchy.
 *
 * 1. For any module, either namespace or folder, one method is defined as
 *    `suite`.
 *
 * 2. The `suite` returns `char*` which has the same meaing as normal test fn.
 *
 * 3. It must return non-NULL if any test in it failed. How it runs tests is
 *    undefined. Typically it runs a flat list of tests.
 */
test_suite_t test_suites[] = {
    {"mlvm/lib", run_lib_suite},
    {"mlvm/sprng", run_sprng_suite},
};

int main() {
  int i;
  int size          = sizeof(test_suites) / sizeof(test_suite_t);
  int suites_failed = 0;

  for (i = 0; i < size; i++) {
    printf("Running suite: %s\n", test_suites[i].name);
    char* result = test_suites[i].fn();

    if (result != 0) {
      suites_failed++;
      printf("\033[1;31mERROR: %s\033[0m\n", result); /* Red */
    } else {
      printf("\033[1;32mALL TESTS PASSED.\033[0m\n"); /* Green */
    }
  }

  printf("Tests run: %d\n", tests_run);
  if (suites_failed)
    printf("\033[1;31mTest suites failed: %d\033[0m\n", suites_failed);
  return suites_failed ? -1 : 0;
}
