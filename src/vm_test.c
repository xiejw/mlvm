#include "testing/testing.h"

#include "vm.h"

#include <string.h>

static char*
test_vm_new()
{
        struct vm_t* vm = vmNew();
        ASSERT_TRUE("vm", vm != NULL);
        vmFree(vm);
        return NULL;
}

static char*
test_vm_exec()
{
        struct vm_t* vm = vmNew();

        int t1, t2, td;
        {
                // create handles.
                struct shape_t* s = vmShapeNew(vm, 2, (int[]){1, 2});

                t1 = vmTensorNew(vm, F32, s);
                t2 = vmTensorNew(vm, F32, s);
                td = vmTensorNew(vm, F32, s);
                ASSERT_TRUE("t1", t1 == 0);
                ASSERT_TRUE("t2", t2 == 1);
                ASSERT_TRUE("td", td == 2);
        }

        {
                // cp data into t1 and t2.
                float32_t  t1_src[2] = {2.34, 5.67};
                float32_t  t2_src[2] = {4.34, 3.67};
                float32_t* data;

                ASSERT_TRUE("err", OK == vmTensorData(vm, t1, (void**)&data));
                memcpy(data, t1_src, 2 * sizeof(float32_t));
                ASSERT_TRUE("err", OK == vmTensorData(vm, t2, (void**)&data));
                memcpy(data, t2_src, 2 * sizeof(float32_t));
        }

        ASSERT_TRUE("err", OK == vmExec(vm, OP_ADD, NULL, td, t1, t2));

        {
                sds_t s = sdsEmpty();
                vmTensorDump(&s, vm, td);
                const char* expected = "<1, 2> f32 [6.680, 9.340]";
                ASSERT_TRUE("dump", 0 == strcmp(s, expected));
                sdsFree(s);
        }
        vmFree(vm);
        return NULL;
}

static char*
test_vm_batch()
{
        struct vm_t* vm = vmNew();

        int t1, t2, td;
        {
                // create handles.
                struct shape_t* s = vmShapeNew(vm, 2, (int[]){1, 2});

                t1 = vmTensorNew(vm, F32, s);
                t2 = vmTensorNew(vm, F32, s);
                td = vmTensorNew(vm, F32, s);
                ASSERT_TRUE("t1", t1 == 0);
                ASSERT_TRUE("t2", t2 == 1);
                ASSERT_TRUE("td", td == 2);
        }

        {
                // cp data into t1 and t2.
                float32_t  t1_src[2] = {2.34, 5.67};
                float32_t  t2_src[2] = {4.34, 3.67};
                float32_t* data;

                ASSERT_TRUE("err", OK == vmTensorData(vm, t1, (void**)&data));
                memcpy(data, t1_src, 2 * sizeof(float32_t));
                ASSERT_TRUE("err", OK == vmTensorData(vm, t2, (void**)&data));
                memcpy(data, t2_src, 2 * sizeof(float32_t));
        }

        const struct oparg_t program[] = {
            // clang-format off
            {OP_ADD, td, t1, t2, /*has_opt*/ 0},
            {OP_MUL, td, td, -1, /*has_opt*/ 1, {.mode = OPT_MODE_F_BIT, .f = 2.0}},
            // clang-format on
        };

        ASSERT_TRUE("err",
                    OK == vmBatch(vm, sizeof(program) / sizeof(struct oparg_t),
                                  program));

        {
                sds_t s = sdsEmpty();
                vmTensorDump(&s, vm, td);
                const char* expected = "<1, 2> f32 [13.360, 18.680]";
                ASSERT_TRUE("dump", 0 == strcmp(s, expected));
                sdsFree(s);
        }
        vmFree(vm);
        return NULL;
}

char*
run_vm_suite()
{
        RUN_TEST(test_vm_new);
        RUN_TEST(test_vm_exec);
        RUN_TEST(test_vm_batch);
        return NULL;
}
