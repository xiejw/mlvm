
#include "testing/testing.h"

#include "object.h"

static char* test_tensor()
{
        struct obj_tensor_t* t = objTensorNew(2, (int[]){2, 3});
        ASSERT_TRUE("no buf", t->buffer == NULL);
        ASSERT_TRUE("no owner", t->owned == 0);
        ASSERT_TRUE("rank", t->rank == 2);
        ASSERT_TRUE("size", t->size == 6);
        ASSERT_TRUE("dim 0", t->dims[0] == 2);
        ASSERT_TRUE("dim 1", t->dims[1] == 3);
        objTensorFree(t);
        return NULL;
}

char* run_vm_object_suite()
{
        RUN_TEST(test_tensor);
        return NULL;
}
