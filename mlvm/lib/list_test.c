#include "mlvm/testing/test.h"

#include "mlvm/lib/list.h"

static char* test_list_empty() {
  list_int_t lt;
  list_init(&lt);

  ASSERT_TRUE("Expect empty.", 0 == list_size(&lt));

  list_deinit(&lt);
  return NULL;
}

static char* test_list_append_and_size() {
  list_int_t lt;
  list_init(&lt);

  list_append(&lt, 123);
  list_append(&lt, 456);
  ASSERT_TRUE("Expect 2 eles.", 2 == list_size(&lt));

  list_deinit(&lt);
  return NULL;
}

static char* test_list_data() {
  list_int_t lt;
  list_init(&lt);

  list_append(&lt, 123);
  list_append(&lt, 456);
  ASSERT_TRUE("Expect 123.", 123 == lt.data[0]);
  ASSERT_TRUE("Expect 456.", 456 == lt.data[1]);

  list_deinit(&lt);
  return NULL;
}

static char* test_list_set() {
  list_int_t lt;
  list_init(&lt);

  list_append(&lt, 123);
  list_append(&lt, 456);
  ASSERT_TRUE("Expect 123.", 123 == lt.data[0]);
  ASSERT_TRUE("Expect 456.", 456 == lt.data[1]);
  list_set(&lt, 1, 789);

  ASSERT_TRUE("Expect 2 eles.", 2 == list_size(&lt));
  ASSERT_TRUE("Expect 789.", 789 == lt.data[1]);

  list_deinit(&lt);
  return NULL;
}

static char* test_list_append_to_grow() {
  int  i;
  int* ptr;

  list_int_t lt;
  list_init(&lt);

  /* Push one to ensure the buffer is allocated.*/
  list_append(&lt, 123);
  ASSERT_TRUE("Expect 1 ele.", 1 == list_size(&lt));
  ptr = lt.data;

  list_append(&lt, 456);
  ASSERT_TRUE("Buffer should be same.", ptr == lt.data);

  for (i = 2; i < 32; i++) list_append(&lt, 456 + i);
  ASSERT_TRUE("Buffer should be growed.", ptr != lt.data);

  for (i = 2; i < 32; i++) ASSERT_TRUE("ele mismatch", 456 + i == lt.data[i]);

  list_deinit(&lt);
  return NULL;
}

char* run_list_test() {
  RUN_TEST(test_list_empty);
  RUN_TEST(test_list_append_and_size);
  RUN_TEST(test_list_data);
  RUN_TEST(test_list_set);
  RUN_TEST(test_list_append_to_grow);
  return NULL;
}
