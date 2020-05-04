#include <stdlib.h>

#include "mlvm/ir/context.h"

#define MLVM_IR_DEFAULT_ERROR_MSG_CAP 256

ir_context_t* ir_context_create() {
  ir_context_t* ctx       = malloc(sizeof(ir_context_t));
  ctx->error_message_cap_ = MLVM_IR_DEFAULT_ERROR_MSG_CAP;
  ctx->error_message = malloc(MLVM_IR_DEFAULT_ERROR_MSG_CAP * sizeof(char));
  return ctx;
}

void ir_context_free(ir_context_t* ctx) {
  free(ctx->error_message);
  free(ctx);
}
