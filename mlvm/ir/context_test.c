#include "mlvm/testing/test.h"

#include <string.h>

#include "mlvm/ir/context.h"

static char* test_context_error_messages() {
  ir_context_t* ctx = ir_context_create();
  int           err = -1;

  ir_context_set_error(ctx, err, "unexpected error.");

  ASSERT_TRUE("Error message mismatch",
              strcmp("unexpected error.", ctx->error_message) == 0);

  ir_context_set_error(ctx, err, "unexpected error: %d", 123);
  ASSERT_TRUE("Error message mismatch",
              strcmp("unexpected error: 123", ctx->error_message) == 0);

  /* Test reusing the error_message. */
  ir_context_set_error(ctx, err, "unexpected error: %s", ctx->error_message);
  ASSERT_TRUE("Error message mismatch",
              strcmp("unexpected error: unexpected error: 123",
                     ctx->error_message) == 0);

  ir_context_free(ctx);
  return NULL;
}

char* run_context_test() {
  RUN_TEST(test_context_error_messages);
  return NULL;
}
