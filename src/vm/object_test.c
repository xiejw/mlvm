
#include "testing/testing.h"

#include "object.h"

static char* test_gabage_collector_default() {
  int collected;
  ASSERT_TRUE("pool", NULL == obj_tensor_pool);

  collected = objTensorGabageCollector();
  ASSERT_TRUE("pool", NULL == obj_tensor_pool);
  ASSERT_TRUE("count", 0 == collected);
  return NULL;
}

static char* test_gabage_collector_single_item() {
  int           collected;
  obj_tensor_t* t1 = objTensorNew(2, (int[]){1, 2});
  t1->mark         = 0;

  collected = objTensorGabageCollector();
  ASSERT_TRUE("pool", NULL == obj_tensor_pool);
  ASSERT_TRUE("count", 1 == collected);
  return NULL;
}

static char* test_gabage_collector_multiple_items() {
  int           collected;
  obj_tensor_t* t;
  obj_float_t   buf[] = {2.0, 3.0};

  // no buffer.
  objTensorNew(2, (int[]){1, 2});
  // owned buffer.
  t         = objTensorNew(2, (int[]){1, 2});
  t->mark   = 1;
  t->owned  = 1;
  t->buffer = malloc(sizeof(obj_float_t) * 2);
  // not owned.
  objTensorNew(2, (int[]){1, 2})->buffer = buf;

  collected = objTensorGabageCollector();
  ASSERT_TRUE("pool", NULL != obj_tensor_pool);
  ASSERT_TRUE("count", 2 == collected);

  // the second item is marked as 0 in last sweep.
  t->mark   = 1;
  collected = objTensorGabageCollector();
  ASSERT_TRUE("pool", NULL != obj_tensor_pool);
  ASSERT_TRUE("count", 0 == collected);

  // the second item is marked as 0 in last sweep.
  collected = objTensorGabageCollector();
  ASSERT_TRUE("pool", NULL == obj_tensor_pool);
  ASSERT_TRUE("count", 1 == collected);
  return NULL;
}

// static char* test_obj_int() {
//   obj_t* o = objNewInt(123);
//   ASSERT_TRUE("value", 123 == objInt(o));
//   ASSERT_TRUE("kind", objIsInt(o));
//   ASSERT_TRUE("kind", !objIsShape(o));
//   objDecrRefCount(o);
//   return NULL;
// }
//
// static char* test_obj_shape() {
//   obj_t* o = objNewShape(3, (int[]){1, 2, 3});
//   ASSERT_TRUE("rank", 3 == objShape(o)->rank);
//   ASSERT_TRUE("dim0", 1 == objShape(o)->dims[0]);
//   ASSERT_TRUE("dim1", 2 == objShape(o)->dims[1]);
//   ASSERT_TRUE("dim2", 3 == objShape(o)->dims[2]);
//   objDecrRefCount(o);
//   return NULL;
// }
//
// static char* test_obj_array() {
//   obj_t* o = objNewArray(3, (obj_float_t[]){1, 2, 3});
//   ASSERT_TRUE("size", 3 == objArray(o)->size);
//   ASSERT_TRUE("v0", 1 == objArray(o)->value[0]);
//   ASSERT_TRUE("v1", 2 == objArray(o)->value[1]);
//   ASSERT_TRUE("v2", 3 == objArray(o)->value[2]);
//   objDecrRefCount(o);
//   return NULL;
// }

char* run_vm_object_suite() {
  RUN_TEST(test_gabage_collector_default);
  RUN_TEST(test_gabage_collector_single_item);
  RUN_TEST(test_gabage_collector_multiple_items);
  return NULL;
}
