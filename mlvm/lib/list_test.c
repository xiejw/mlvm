#include "mlvm/testing/test.h"

#include "mlvm/lib/list.h"

static char* emtpy_list() {
  list_int_t lt;
  list_init(&lt);

  ASSERT_TRUE("Expect empty.", 0 == list_size(&lt));
  /*
  list_append(&lt, 123);
  list_append(&lt, 456);
  printf("List 0 %d\n", list_get(&lt, 0));
  printf("List 1 %d\n", list_get(&lt, 1));
  printf("List size %lld\n", list_size(&lt));
  */

  list_deinit(&lt);
  return NULL;
}

char* run_list_test() {
  RUN_TEST(emtpy_list);
  return NULL;
}
