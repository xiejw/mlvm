#include "testing/testing.h"

#include <string.h>

#include "adt/sds.h"

#include "object.h"

static char* test_tensor()
{
        struct obj_tensor_t* t =
            objTensorNew(OBJ_DTYPE_FLOAT32, 2, (int[]){2, 3});
        ASSERT_TRUE("no buf", t->buffer == NULL);
        ASSERT_TRUE("no owner", t->owned == 0);
        ASSERT_TRUE("rank", t->rank == 2);
        ASSERT_TRUE("dtype", t->dtype == OBJ_DTYPE_FLOAT32);
        ASSERT_TRUE("size", t->size == 6);
        ASSERT_TRUE("dim 0", t->dims[0] == 2);
        ASSERT_TRUE("dim 1", t->dims[1] == 3);

        {
                sds_t s = sdsEmpty();
                objTensorDump(t, &s);
                sdsFree(s);
        }

        objTensorFree(t);
        return NULL;
}

static char* test_tensor_dump_float()
{
        sds_t                s = sdsEmpty();
        struct obj_tensor_t* t = objTensorNewFloat32(/*rank=2*/ 2, 2, 3);

        objTensorDump(t, &s);
        ASSERT_TRUE("dump", 0 == strcmp("[ (NULL) ]", s));

        sdsFree(s);
        objTensorFree(t);
        return NULL;
}

char* run_vm_object_suite()
{
        RUN_TEST(test_tensor);
        RUN_TEST(test_tensor_dump_float);
        return NULL;
}
