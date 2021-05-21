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

#define CHECK_TENSOR(vm, td, expect, fmt, ...)         \
        if (expect_dump(vm, td, expect)) {             \
                errDump(fmt, __VA_ARGS__);             \
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
        struct shape_t* s  = vmShapeNew(vm, 2, (int[]){1, 3});

        int t1 = vmTensorNew(vm, F32, s);
        int t2 = vmTensorNew(vm, F32, s);
        int td = vmTensorNew(vm, F32, s);

        COPY_DATA(vm, t1, 3, ((float32_t[]){2.34, 5.67, 2.00}));
        COPY_DATA(vm, t2, 3, ((float32_t[]){4.34, 3.67, 2.00}));

        enum opcode_t ops[] = {OP_ADD, OP_MUL, OP_MINUS,
                               OP_MAX, OP_EQ,  OP_CMPL};

        const char* expected_strs[] = {
            "<1, 3> f32 [6.680, 9.340, 4.000]",
            "<1, 3> f32 [10.156, 20.809, 4.000]",
            "<1, 3> f32 [-2.000, 2.000, 0.000]",
            "<1, 3> f32 [4.340, 5.670, 2.000]",
            "<1, 3> f32 [0.000, 0.000, 1.000]",
            "<1, 3> f32 [0.000, 1.000, 0.000]",
        };

        for (int i = 0; i < sizeof(ops) / sizeof(enum opcode_t); i++) {
                NE(vmExec(vm, ops[i], NULL, td, t1, t2));
                CHECK_TENSOR(vm, td, expected_strs[i], "failed at %d-th Op\n",
                             i);
        }
        vmFree(vm);
        return NULL;
}

static char*
test_element_ops_with_scalar_rhs()
{
        struct vm_t*    vm = vmNew();
        struct shape_t* s  = vmShapeNew(vm, 2, (int[]){1, 3});
        struct shape_t* s1 = vmShapeNew(vm, 1, (int[]){1});

        int t1 = vmTensorNew(vm, F32, s);
        int t2 = vmTensorNew(vm, F32, s1);
        int td = vmTensorNew(vm, F32, s);

        COPY_DATA(vm, t1, 3, ((float32_t[]){2.34, 5.67, 3.67}));
        COPY_DATA(vm, t2, 1, ((float32_t[]){3.67}));

        enum opcode_t ops[] = {OP_ADD, OP_MUL, OP_MINUS,
                               OP_MAX, OP_EQ,  OP_CMPL};

        const char* expected_strs[] = {
            "<1, 3> f32 [6.010, 9.340, 7.340]",
            "<1, 3> f32 [8.588, 20.809, 13.469]",
            "<1, 3> f32 [-1.330, 2.000, 0.000]",
            "<1, 3> f32 [3.670, 5.670, 3.670]",
            "<1, 3> f32 [0.000, 0.000, 1.000]",
            "<1, 3> f32 [0.000, 1.000, 0.000]",
        };

        for (int i = 0; i < sizeof(ops) / sizeof(enum opcode_t); i++) {
                NE(vmExec(vm, ops[i], NULL, td, t1, t2));
                CHECK_TENSOR(vm, td, expected_strs[i], "failed at %d-th Op\n",
                             i);
        }
        vmFree(vm);
        return NULL;
}

static char*
test_element_ops_f()
{
        struct vm_t*    vm = vmNew();
        struct shape_t* s  = vmShapeNew(vm, 2, (int[]){1, 3});

        int t1 = vmTensorNew(vm, F32, s);
        int td = vmTensorNew(vm, F32, s);

        COPY_DATA(vm, t1, 3, ((float32_t[]){2.34, 5.67, 3.000}));

        enum opcode_t ops[] = {OP_ADD, OP_MUL, OP_MINUS,
                               OP_MAX, OP_EQ,  OP_CMPL};

        const char* expected_strs[] = {
            "<1, 3> f32 [5.340, 8.670, 6.000]",
            "<1, 3> f32 [7.020, 17.010, 9.000]",
            "<1, 3> f32 [-0.660, 2.670, 0.000]",
            "<1, 3> f32 [3.000, 5.670, 3.000]",
            "<1, 3> f32 [0.000, 0.000, 1.000]",
            "<1, 3> f32 [0.000, 1.000, 0.000]",
        };

        const struct opopt_t opt = {.mode = OPT_MODE_F_BIT, .f = 3};
        for (int i = 0; i < sizeof(ops) / sizeof(enum opcode_t); i++) {
                NE(vmExec(vm, ops[i], &opt, td, t1, -1));
                CHECK_TENSOR(vm, td, expected_strs[i], "failed at %d-th Op\n",
                             i);
        }
        vmFree(vm);
        return NULL;
}

static char*
test_matmul()
{
        struct vm_t*    vm = vmNew();
        struct shape_t* s  = vmShapeNew(vm, 2, (int[]){2, 2});

        int t1 = vmTensorNew(vm, F32, s);
        int t2 = vmTensorNew(vm, F32, s);
        int td = vmTensorNew(vm, F32, s);

        COPY_DATA(vm, t1, 4, ((float32_t[]){2.34, 5.67, -1.23, 2.34}));
        COPY_DATA(vm, t2, 4, ((float32_t[]){4.34, 3.67, -2.24, 3.45}));

        struct opopt_t opt1 = {.mode = OPT_MATMUL_TRANS_NOT};
        struct opopt_t opt2 = {.mode = OPT_MATMUL_TRANS_LHS};
        struct opopt_t opt3 = {.mode = OPT_MATMUL_TRANS_RHS};

        struct opopt_t* opts[] = {NULL, &opt1, &opt2, &opt3};

        const char* expected_strs[] = {
            "<2, 2> f32 [-2.545, 28.149, -10.580, 3.559]",
            "<2, 2> f32 [-2.545, 28.149, -10.580, 3.559]",
            "<2, 2> f32 [12.911, 4.344, 19.366, 28.882]",
            "<2, 2> f32 [30.965, 14.320, 3.250, 10.828]",
        };

        for (int i = 0; i < sizeof(opts) / sizeof(struct opopt_t*); i++) {
                NE(vmExec(vm, OP_MATMUL, opts[i], td, t1, t2));
                CHECK_TENSOR(vm, td, expected_strs[i], "failed at %d-th opt\n",
                             i);
        }

        vmFree(vm);
        return NULL;
}

char*
run_op_suite()
{
        RUN_TEST(test_element_ops);
        RUN_TEST(test_element_ops_with_scalar_rhs);
        RUN_TEST(test_element_ops_f);
        RUN_TEST(test_matmul);
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
                return errNew("info:\nexpected: %s\ngot     : %s\n", expected,
                              s);
        }
        sdsFree(s);
        return OK;
}
