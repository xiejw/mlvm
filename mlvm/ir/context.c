#include <stdlib.h>

#include <assert.h>
#include <stdarg.h>
#include <stdio.h>

#include "mlvm/ir/context.h"

#define MLVM_IR_DEFAULT_ERROR_MSG_CAP 256

ir_context_t* ir_context_create() {
  ir_context_t* ctx       = malloc(sizeof(ir_context_t));
  ctx->prng               = NULL;
  ctx->error_message_cap_ = MLVM_IR_DEFAULT_ERROR_MSG_CAP;
  ctx->error_message = malloc(MLVM_IR_DEFAULT_ERROR_MSG_CAP * sizeof(char));
  ctx->error_message_buffer_ =
      malloc(MLVM_IR_DEFAULT_ERROR_MSG_CAP * sizeof(char));
  return ctx;
}

void ir_context_free(ir_context_t* ctx) {
  if (ctx->prng != NULL) sprng_free(ctx->prng);
  free(ctx->error_message);
  free(ctx->error_message_buffer_);
  free(ctx);
}

void ir_context_set_prng(ir_context_t* ctx, sprng_t* prng) {
  if (ctx->prng != NULL) sprng_free(ctx->prng);
  ctx->prng = prng;
}

int ir_context_set_error(ir_context_t* ctx, int error_code, char* fmt, ...) {
  char* output_buffer = ctx->error_message_buffer_;
  assert(error_code < 0);

  {
    va_list args;
    int     n;
    va_start(args, fmt);
    output_buffer = malloc(MLVM_IR_DEFAULT_ERROR_MSG_CAP * sizeof(char));
    n = vsnprintf(output_buffer, MLVM_IR_DEFAULT_ERROR_MSG_CAP, fmt, args);
    assert(n < MLVM_IR_DEFAULT_ERROR_MSG_CAP - 1);
    va_end(args);
  }

  /* swapping the buffers now. */
  ctx->error_message_buffer_ = ctx->error_message;
  ctx->error_message         = output_buffer;
  return error_code;
}
