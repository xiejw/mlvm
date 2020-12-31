
#include "testing/testing.h"

#include "object.h"

static char* test_obj_int64() {
  obj_t* o = objNewInt64(123);
  ASSERT_TRUE("value", 123 == objInt64Value(o));
  objDecrRefCount(o);
  return NULL;
}

static char* test_obj_shape() {
  int    dims[] = {1, 2, 3};
  obj_t* o      = objNewShape(3, dims);
  ASSERT_TRUE("rank", 3 == objShapeRank(o));
  ASSERT_TRUE("dim0", 1 == objShapeDims(o)[0]);
  ASSERT_TRUE("dim1", 2 == objShapeDims(o)[1]);
  ASSERT_TRUE("dim2", 3 == objShapeDims(o)[2]);
  objDecrRefCount(o);
  return NULL;
}

char* run_vm_object_suite() {
  RUN_TEST(test_obj_int64);
  RUN_TEST(test_obj_shape);
  return NULL;
}
