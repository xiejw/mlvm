#ifndef MLVM_TESTING_TEST_H_
#define MLVM_TESTING_TEST_H_

#include <math.h>  /* fabs */
#include <stdio.h> /* printf */

#ifdef DEBUG_TEST_ASSERT
#define DEBUG_TEST_PRINTF(x, ...) printf(x, __VA_ARGS__)
#else
#define DEBUG_TEST_PRINTF(x, ...)
#endif

/* This unit testing framework is inspired by MinUnit. */

extern int tests_run;

#define RUN_TEST(test) RUN_TEST_IMPL(test, #test)

#define RUN_TEST_IMPL(test, name)  \
  do {                             \
    printf("  Running: %s", name); \
    char *msg = (test)();          \
    tests_run++;                   \
    if (msg != NULL) {             \
      printf("...Failed.\n");      \
      return msg;                  \
    } else {                       \
      printf(".\n");               \
    }                              \
  } while (0)

/* RUN_SUITE does not increase test count. */
#define RUN_SUITE(test)          \
  do {                           \
    char *msg = (test)();        \
    if (msg != NULL) return msg; \
  } while (0)

#define ASSERT_TRUE(msg, test) \
  do {                         \
    if (!(test)) return msg;   \
  } while (0)

#define ASSERT_ARRAY_CLOSE(msg, expected, got, size, tol)                     \
  do {                                                                        \
    int i;                                                                    \
    for (i = 0; i < size; i++) {                                              \
      if (fabs(expected[i] - got[i]) >= tol) {                                \
        DEBUG_TEST_PRINTF("\n-> at element %d\n-> got %f\n-> expect %f\n", i, \
                          expected[i], got[i]);                               \
        return msg;                                                           \
      }                                                                       \
    }                                                                         \
  } while (0)

#endif
