#ifndef MLVM_IR_CONTEXT_H_
#define MLVM_IR_CONTEXT_H_

#include "mlvm/sprng/sprng.h"

typedef struct {
  char*    error_message;
  sprng_t* prng; /* Optional. */

  /* Internal fields. */
  unsigned int error_message_cap_;
} ir_context_t;

extern ir_context_t* ir_context_create();
extern void          ir_context_free();

/* Move the prng into the context. Currrent `prng` is freed. */
extern void ir_context_set_prng(ir_context_t* ctx, sprng_t* prng);

#endif
