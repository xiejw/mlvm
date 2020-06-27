#include "mlvm/testing/test.h"

#include "mlvm/runtime/kernel/kernel.h"

static char* test_add() {
  double value_1[] = {1.0, 2.0};
  double value_2[] = {11.0, 12.0};
  double output_value[2];
  double expected[2] = {12, 14};

  mlvm_uint_t shape_1x2[] = {1, 2};

  tensor_t* t_1 = tensor_create(2, shape_1x2, value_1, MLVM_ALIAS_VALUE);
  tensor_t* t_2 = tensor_create(2, shape_1x2, value_2, MLVM_ALIAS_VALUE);
  tensor_t* t_output =
      tensor_create(2, shape_1x2, output_value, MLVM_ALIAS_VALUE);

  kernel_add(t_output, t_1, t_2);
  ASSERT_ARRAY_CLOSE("Result mismatch", expected, t_output->value, 2, 1e-6);

  tensor_free(t_1);
  tensor_free(t_2);
  tensor_free(t_output);
  return NULL;
}

char* run_add_test() {
  RUN_TEST(test_add);
  return NULL;
}
