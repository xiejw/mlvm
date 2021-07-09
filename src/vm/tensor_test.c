#include "testing/testing.h"

#include "vm.h"

#include <string.h>

static char *
test_tensor_new()
{
        struct vm_t *vm = vmNew();

        for (int i = 0; i < 10; i++) {
                struct shape_t *s  = vmShapeNew(vm, 2, (int[]){3, 4});
                int             td = vmTensorNew(vm, F32, s);
                ASSERT_TRUE("td", td == i);
        }
        for (int i = 0; i < 10; i++) {
                ASSERT_TRUE("err", OK == vmTensorFree(vm, i));
        }
        vmFree(vm);
        return NULL;
}

static char *
test_tensor_info()
{
        struct vm_t *vm = vmNew();

        int td;
        {
                struct shape_t *s = vmShapeNew(vm, 2, (int[]){3, 4});
                td                = vmTensorNew(vm, F32, s);
                ASSERT_TRUE("td", td == 0);
        }

        {
                enum data_t     dtype;
                struct shape_t *s1;
                ASSERT_TRUE("err", OK == vmTensorInfo(vm, td, &dtype, &s1));
                ASSERT_TRUE("dtype", dtype == F32);
                ASSERT_TRUE("rank", s1->rank == 2);
                ASSERT_TRUE("dim 0", s1->dims[0] == 3);
                ASSERT_TRUE("dim 1", s1->dims[1] == 4);
                ASSERT_TRUE("size", s1->size == 12);
        }

        vmFree(vm);
        return NULL;
}

static char *
test_tensor_data()
{
        struct vm_t *vm = vmNew();

        int td;
        {
                struct shape_t *s = vmShapeNew(vm, 2, (int[]){1, 2});
                td                = vmTensorNew(vm, F32, s);
                ASSERT_TRUE("td", td == 0);
        }

        {
                f32_t  src[2] = {2.34, 5.67};
                f32_t *data;
                ASSERT_TRUE("err", OK == vmTensorData(vm, td, (void **)&data));
                memcpy(data, src, 2 * sizeof(f32_t));
        }

        {
                sds_t s = sdsEmpty();
                vmTensorDump(&s, vm, td);
                const char *expected = "<1, 2> f32 [2.340, 5.670]";
                ASSERT_TRUE("dump", 0 == strcmp(s, expected));
                sdsFree(s);
        }
        vmFree(vm);
        return NULL;
}

static char *
test_tensor_swap()
{
        struct vm_t *vm = vmNew();

        int td;
        {
                struct shape_t *s = vmShapeNew(vm, 2, (int[]){1, 2});
                td                = vmTensorNew(vm, F32, s);
                ASSERT_TRUE("td", td == 0);
        }

        {
                f32_t *data = malloc(2 * sizeof(f32_t));
                data[0]     = 2.34;
                data[1]     = 5.67;
                ASSERT_TRUE("err", OK == vmTensorSwap(vm, td, (void **)&data));
                free(data);
        }

        {
                sds_t s = sdsEmpty();
                vmTensorDump(&s, vm, td);
                const char *expected = "<1, 2> f32 [2.340, 5.670]";
                ASSERT_TRUE("dump", 0 == strcmp(s, expected));
                sdsFree(s);
        }
        vmFree(vm);
        return NULL;
}

char *
run_vm_tensor_suite()
{
        RUN_TEST(test_tensor_new);
        RUN_TEST(test_tensor_info);
        RUN_TEST(test_tensor_data);
        RUN_TEST(test_tensor_swap);
        return NULL;
}
