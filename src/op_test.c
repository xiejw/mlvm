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

#define CHECK_TENSOR(vm, td, expect)                   \
        if (expect_dump(vm, td, expect)) {             \
                errDump("tensor mismatch.\n"); \
                return "failed to expect tensor dump"; \
        }

#define NE(e)                               \
        if ((e)) {                          \
                return "unexpected error."; \
        }

// -----------------------------------------------------------------------------
// unit tests.
// -----------------------------------------------------------------------------

static char*
test_element_ops()
{
        struct vm_t*    vm = vmNew();
        struct shape_t* s  = vmShapeNew(vm, 2, (int[]){1, 2});

        int t1 = vmTensorNew(vm, F32, s);
        int t2 = vmTensorNew(vm, F32, s);
        int td = vmTensorNew(vm, F32, s);

        COPY_DATA(vm, t1, 2, ((float32_t[]){2.34, 5.67}));
        COPY_DATA(vm, t2, 2, ((float32_t[]){4.34, 3.67}));

        enum opcode_t ops[]           = {OP_ADD, OP_MUL, OP_MINUS, OP_MAX};
        const char*   expected_strs[] = {
            "<1, 2> f32 [6.680, 9.340]",
            "<1, 2> f32 [6.680, 9.340]",
            "<1, 2> f32 [-2.000, 2.000]",
            "<1, 2> f32 [4.340, 5.670]",
        };

        for (int i = 0; i < sizeof(ops) / sizeof(enum opcode_t); i++) {
                NE(vmExec(vm, ops[i], NULL, td, t1, t2));
                CHECK_TENSOR(vm, td, expected_strs[i]);
        }
        vmFree(vm);
        return NULL;
}

char*
run_op_suite()
{
        RUN_TEST(test_element_ops);
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
                return errNew("info:\nexpected: %s\ngot     : %s\n", expected, s);
        }
        sdsFree(s);
        return OK;
}
