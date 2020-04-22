#ifndef MLVM_TESTING_TEST_H_
#define MLVM_TESTING_TEST_H_

/* This unit testing framework is inspired by MinUnit. */

extern int tests_run;

#define RUN_TEST(test)           \
  do {                           \
    char *msg = (test)();        \
    tests_run++;                 \
    if (msg != NULL) return msg; \
  } while (0)

#define ASSERT_TRUE(msg, test) \
  do {                         \
    if (!(test)) return msg;   \
  } while (0)

#endif
