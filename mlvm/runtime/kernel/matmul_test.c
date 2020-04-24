#include "mlvm/testing/test.h"

#include "mlvm/runtime/kernel/kernel.h"

static char* test_matmul() {
  double value_1[] = {1.0, 2.0, 3.0, 4.0, 5.0, 6.0};
  double value_2[] = {11.0, 12.0, 13.0, 14.0, 15.0, 16.0};
  double output_value[4];
  double expected[4] = {82, 88, 199, 214};

  tensor_shape_t shape_2x3[] = {2, 3};
  tensor_shape_t shape_3x2[] = {3, 2};
  tensor_shape_t shape_2x2[] = {2, 2};

  tensor_t* t_1 = tensor_create(2, shape_2x3, value_1, MLVM_ALIAS_VALUE);
  tensor_t* t_2 = tensor_create(2, shape_3x2, value_2, MLVM_ALIAS_VALUE);
  tensor_t* t_output =
      tensor_create(2, shape_2x2, output_value, MLVM_ALIAS_VALUE);

  kernel_matmul(t_output, t_1, t_2);
  ASSERT_ARRAY_CLOSE("Result mismatch", expected, t_output->value, 4, 1e-6);

  tensor_free(t_1);
  tensor_free(t_2);
  tensor_free(t_output);
  return NULL;
}

static char* test_matmul_with_arg_2_stride() {
  double        value_1[] = {1.0, 2.0, 3.0, 4.0, 5.0, 6.0};
  double        value_2[] = {11.0, 12.0, 13.0, 14.0, 15.0, 16.0};
  double        output_value[4];
  double        expected[4]  = {74, 92, 182, 227};
  tensor_size_t new_stride[] = {1, 3};

  tensor_shape_t shape_2x3[] = {2, 3};
  tensor_shape_t shape_3x2[] = {3, 2};
  tensor_shape_t shape_2x2[] = {2, 2};

  tensor_t* t_1 = tensor_create(2, shape_2x3, value_1, MLVM_ALIAS_VALUE);
  tensor_t* t_2 = tensor_create(2, shape_3x2, value_2, MLVM_ALIAS_VALUE);
  tensor_t* t_output =
      tensor_create(2, shape_2x2, output_value, MLVM_ALIAS_VALUE);

  tensor_set_stride(t_2, new_stride);
  kernel_matmul(t_output, t_1, t_2);
  ASSERT_ARRAY_CLOSE("Result mismatch", expected, t_output->value, 4, 1e-6);

  tensor_free(t_1);
  tensor_free(t_2);
  tensor_free(t_output);
  return NULL;
}

char* run_matmul_test() {
  RUN_TEST(test_matmul);
  RUN_TEST(test_matmul_with_arg_2_stride);
  return NULL;
}
