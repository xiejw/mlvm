
#include "testing/testing.h"

#include "object.h"

static char* test_obj_int64() {
  obj_t* o = objNewInt64(123);
  ASSERT_TRUE("value", 123 == objInt64V(o));
  objDecrRefCount(o);
  return NULL;
}

static char* test_obj_shape() {
  obj_t* o = objNewShape(3, (int[]){1, 2, 3});
  ASSERT_TRUE("rank", 3 == objShape(o)->rank);
  ASSERT_TRUE("dim0", 1 == objShape(o)->dims[0]);
  ASSERT_TRUE("dim1", 2 == objShape(o)->dims[1]);
  ASSERT_TRUE("dim2", 3 == objShape(o)->dims[2]);
  objDecrRefCount(o);
  return NULL;
}

char* run_vm_object_suite() {
  RUN_TEST(test_obj_int64);
  RUN_TEST(test_obj_shape);
  return NULL;
}
