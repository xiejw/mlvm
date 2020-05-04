#ifndef MLVM_IR_CONTEXT_H_
#define MLVM_IR_CONTEXT_H_

typedef struct {
  char* error_message;

  /* Internal fields. */
  unsigned int error_message_cap_;
} ir_context_t;

extern ir_context_t* ir_context_create();
extern void          ir_context_free();

#endif
