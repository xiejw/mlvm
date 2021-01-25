
// #include "testing/testing.h"
//
// #include "object.h"
//
// static char* test_gc()
// {
//         int collected;
//         ASSERT_TRUE("pool", NULL == obj_tensor_pool);
//
//         collected = objGC();
//         ASSERT_TRUE("pool", NULL == obj_tensor_pool);
//         ASSERT_TRUE("count", 0 == collected);
//         return NULL;
// }
//
// static char* test_gc_single_item()
// {
//         int                  collected;
//         struct obj_tensor_t* t1 = objTensorNew(2, (int[]){1, 2});
//         t1->mark                = 0;
//
//         collected = objGC();
//         ASSERT_TRUE("pool", NULL == obj_tensor_pool);
//         ASSERT_TRUE("count", 1 == collected);
//         return NULL;
// }
//
// static char* test_gc_multiple_items()
// {
//         int                  collected;
//         struct obj_tensor_t* t;
//         obj_float_t          buf[] = {2.0, 3.0};
//
//         // no buffer.
//         objTensorNew(2, (int[]){1, 2});
//         // owned buffer.
//         t         = objTensorNew(2, (int[]){1, 2});
//         t->mark   = 1;
//         t->owned  = 1;
//         t->buffer = malloc(sizeof(obj_float_t) * 2);
//         // not owned.
//         objTensorNew(2, (int[]){1, 2})->buffer = buf;
//
//         collected = objGC();
//         ASSERT_TRUE("pool", NULL != obj_tensor_pool);
//         ASSERT_TRUE("count", 2 == collected);
//
//         // the second item is marked as 0 in last sweep.
//         t->mark   = 1;
//         collected = objGC();
//         ASSERT_TRUE("pool", NULL != obj_tensor_pool);
//         ASSERT_TRUE("count", 0 == collected);
//
//         // the second item is marked as 0 in last sweep.
//         collected = objGC();
//         ASSERT_TRUE("pool", NULL == obj_tensor_pool);
//         ASSERT_TRUE("count", 1 == collected);
//         return NULL;
// }
//
// char* run_vm_object_suite()
// {
//         RUN_TEST(test_gc);
//         RUN_TEST(test_gc_single_item);
//         RUN_TEST(test_gc_multiple_items);
//         return NULL;
// }
