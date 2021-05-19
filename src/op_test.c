#include "testing/testing.h"

#include "vm.h"

#include <string.h>

// -----------------------------------------------------------------------------
// helper prototype
// -----------------------------------------------------------------------------
error_t copy_data(struct vm_t* vm, int td, size_t size, float32_t* src);
error_t expect_dump(struct vm_t* vm, int td, const char* expected);

#define COPY_DATA(vm, td, size, src)          \
        if (copy_data(vm, td, size, src)) {   \
                return "failed to copy data"; \
        }

#define EXPECT_DUMP(vm, td, expect)                    \
        if (expect_dump(vm, td, expect)) {             \
                return "failed to expect tensor dump"; \
        }

static char*
test_op_add()
{
        struct vm_t*    vm = vmNew();
        struct shape_t* s  = vmShapeNew(vm, 2, (int[]){1, 2});

        int t1 = vmTensorNew(vm, F32, s);
        int t2 = vmTensorNew(vm, F32, s);
        int td = vmTensorNew(vm, F32, s);

        COPY_DATA(vm, t1, 2, ((float32_t[]){2.34, 5.67}));
        COPY_DATA(vm, t2, 2, ((float32_t[]){4.34, 3.67}));

        ASSERT_TRUE("err", OK == vmExec(vm, OP_ADD, NULL, td, t1, t2));
        EXPECT_DUMP(vm, td, "<1, 2> f32 [6.680, 9.340]");
        vmFree(vm);
        return NULL;
}

char*
run_op_suite()
{
        RUN_TEST(test_op_add);
        return NULL;
}

// -----------------------------------------------------------------------------
// helper impl
// -----------------------------------------------------------------------------
error_t
copy_data(struct vm_t* vm, int td, size_t size, float32_t* src)
{
        float32_t* data;
        error_t    err = vmTensorData(vm, td, (void**)&data);
        if (err) return err;
        memcpy(data, src, size * sizeof(float32_t));
        return OK;
}

error_t
expect_dump(struct vm_t* vm, int td, const char* expected)
{
        sds_t s = sdsEmpty();
        vmTensorDump(&s, vm, td);
        if (0 != strcmp(s, expected)) {
                sdsFree(s);
                return errNew("mismatch.");
        }
        sdsFree(s);
        return OK;
}
