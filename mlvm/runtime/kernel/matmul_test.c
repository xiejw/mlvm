#include "mlvm/testing/test.h"

#include "mlvm/runtime/kernel/kernel.h"

static char* test_matmul() {
  double value_1[] = {1.0, 2.0, 3.0, 4.0, 5.0, 6.0};
  double value_2[] = {11.0, 12.0, 13.0, 14.0, 15.0, 16.0};
  double output_value[4];
  double expected[4] = {82, 88, 199, 214};

  tensor_t* t_1 =
      tensor_create(2, (tensor_shape_t[]){2, 3}, value_1, MLVM_ALIAS_VALUE);
  tensor_t* t_2 =
      tensor_create(2, (tensor_shape_t[]){3, 2}, value_2, MLVM_ALIAS_VALUE);
  tensor_t* t_output = tensor_create(2, (tensor_shape_t[]){2, 2}, output_value,
                                     MLVM_ALIAS_VALUE);

  ASSERT_TRUE("Should be succesful", 0 == kernel_matmul(t_output, t_1, t_2));
  ASSERT_ARRAY_CLOSE("Result mismatch", expected, t_output->value, 4, 1e-6);

  tensor_free(t_1);
  tensor_free(t_2);
  tensor_free(t_output);
  return NULL;
}

char* run_matmul_test() {
  RUN_TEST(test_matmul);
  return NULL;
}
