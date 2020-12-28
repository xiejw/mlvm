
#include "testing/testing.h"

#include "object.h"

static char* test_obj_int64() {
  obj_t* o = objNewInt64(123);
  ASSERT_TRUE("value", 123 == objInt64Value(o));
  objDecrRefCount(o);
  return NULL;
}

char* run_vm_object_suite() {
  RUN_TEST(test_obj_int64);
  return NULL;
}
