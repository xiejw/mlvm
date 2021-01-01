
#include "testing/testing.h"

#include "object.h"

static char* test_obj_int() {
  obj_t* o = objNewInt(123);
  ASSERT_TRUE("value", 123 == objInt(o));
  ASSERT_TRUE("kind", objIsInt(o));
  ASSERT_TRUE("kind", !objIsShape(o));
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

static char* test_obj_array() {
  obj_t* o = objNewArray(3, (obj_float_t[]){1, 2, 3});
  ASSERT_TRUE("size", 3 == objArray(o)->size);
  ASSERT_TRUE("v0", 1 == objArray(o)->value[0]);
  ASSERT_TRUE("v1", 2 == objArray(o)->value[1]);
  ASSERT_TRUE("v2", 3 == objArray(o)->value[2]);
  objDecrRefCount(o);
  return NULL;
}

char* run_vm_object_suite() {
  RUN_TEST(test_obj_int);
  RUN_TEST(test_obj_shape);
  RUN_TEST(test_obj_array);
  return NULL;
}
