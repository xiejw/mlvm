#ifndef MLVM_IR_FUNCTION_H_
#define MLVM_IR_FUNCTION_H_

#include <stdint.h>

#include "mlvm/ir/tensor.h"
#include "mlvm/lib/list.h"

/* For now, we use static size. */
#define MLVM_IR_MAX_OPERANDS_PER_INS 4
#define MLVM_IR_MAX_OUTPUTS_PER_INS  10

typedef enum { IR_CONST } ir_operand_type;

typedef union {
  tensor_t* tensor; /* Owned the tensor. */
} ir_operand_value;

typedef struct {
  char*            name;
  ir_operand_type  type;
  ir_operand_value value;
} ir_operand_t;

typedef enum { IR_OP_ADD } ir_instruction_type;

typedef struct {
  char*               name;
  ir_instruction_type type;
  ir_operand_t*       operands[MLVM_IR_MAX_OPERANDS_PER_INS];
  ir_operand_t*       outputs[MLVM_IR_MAX_OUTPUTS_PER_INS];
} ir_instrution;

typedef list_t(ir_operand_t*) list_ir_operand_t;

typedef struct {
  char*             name; /* Function name. */
  list_ir_operand_t const_tensors;
} ir_function_t;

extern ir_function_t* ir_function_create(char* name);
extern void           ir_function_free(ir_function_t* func);
extern int            ir_function_print(ir_function_t* func, int fd);

/*
 * Args:
 *     value_mode can only be MLVM_COPY_VALUE, MLVM_ALIAS_VALUE, or
 *     MLVM_MOVE_VALUE. For MLVM_MOVE_VALUE, the original tensor is invalid for
 *     usage after the invocation. In this case, `tensor` must own the value.
 *
 * Returns:
 *     NULL for error. The returned operand is owned by the function.
 */
extern ir_operand_t* ir_function_add_constant(ir_function_t* func,
                                              tensor_t* tensor, int value_mode);

#endif
