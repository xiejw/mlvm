#include "testing/testing.h"

#include <string.h>

#include "adt/sds.h"

#include "tensor.h"

static char* test_new()
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

static char* test_dump_null()
{
        sds_t                s = sdsEmpty();
        struct obj_tensor_t* t = objTensorNewFloat32(/*rank=2*/ 2, 2, 3);

        objTensorDump(t, &s);
        ASSERT_TRUE("dump", 0 == strcmp("[ (NULL) ]", s));

        sdsFree(s);
        objTensorFree(t);
        return NULL;
}

static char* test_dump_float32()
{
        sds_t                s = sdsEmpty();
        struct obj_tensor_t* t = objTensorNewFloat32(/*rank=*/1, 2);

        objTensorAllocAndCopy(t, (float[]){1.2, 2.3});

        objTensorDump(t, &s);
        ASSERT_TRUE("dump", 0 == strcmp("[ 1.200000, 2.300000,]", s) ||
                                (printf("\ngot: %s\n", s), 0));

        sdsFree(s);
        objTensorFree(t);
        return NULL;
}

static char* test_dump_int32()
{
        sds_t                s = sdsEmpty();
        struct obj_tensor_t* t = objTensorNewInt32(/*rank=*/1, 2);

        objTensorAllocAndCopy(t, (int32_t[]){1, 2});

        objTensorDump(t, &s);
        ASSERT_TRUE("dump", 0 == strcmp("[ 1, 2,]", s) ||
                                (printf("\ngot: %s\n", s), 0));

        sdsFree(s);
        objTensorFree(t);
        return NULL;
}

char* run_vm_tensor_suite()
{
        RUN_TEST(test_new);
        RUN_TEST(test_dump_null);
        RUN_TEST(test_dump_float32);
        RUN_TEST(test_dump_int32);
        return NULL;
}
