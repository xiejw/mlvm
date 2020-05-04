#include "mlvm/testing/test.h"

#include "mlvm/ir/context.h"

static char* test_context_error_messages() {
  ir_context_t* ctx = ir_context_create();
  ir_context_free(ctx);
  return NULL;
}

char* run_context_test() {
  RUN_TEST(test_context_error_messages);
  return NULL;
}
