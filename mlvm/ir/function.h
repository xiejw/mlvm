#ifndef MLVM_IR_FUNCTION_H_
#define MLVM_IR_FUNCTION_H_

#include <stdint.h>

#include "mlvm/ir/tensor.h"
#include "mlvm/lib/list.h"

/******************************************************************************
 * Operands.
 *
 * operrand.c
 *****************************************************************************/

typedef enum { IR_OPERAND_CONST } ir_operand_type;

typedef union {
  tensor_t* tensor; /* Owned the tensor. */
} ir_operand_value;

typedef struct {
  char*            name;
  ir_operand_type  type;
  ir_operand_value value;
} ir_operand_t;

typedef list_t(ir_operand_t*) list_ir_operand_t;

extern ir_operand_t* ir_operand_create_const(tensor_t*   const_tensor,
                                             const char* name_fmt, ...);
extern void          ir_operand_free(ir_operand_t* operand);

/******************************************************************************
 * Instruction.
 *
 * instruction.c
 *****************************************************************************/

struct ir_function_t;

typedef enum { IR_OP_ADD } ir_instruction_type;

typedef struct {
  char*                 name;
  ir_instruction_type   type;
  struct ir_function_t* parent_func;

  list_ir_operand_t operands;
  list_ir_operand_t outputs;
} ir_instruction_t;

extern ir_instruction_t* ir_instruction_create(
    struct ir_function_t* parent_func, char* name, ir_instruction_type type);
/* extern void ir_instru
 */

/******************************************************************************
 * Function.
 *****************************************************************************/

typedef struct ir_function_t {
  char*             name; /* Function name. */
  list_ir_operand_t const_tensors;
} ir_function_t;

extern ir_function_t* ir_function_create(char* name);
extern void           ir_function_free(ir_function_t* func);
extern int            ir_function_print(ir_function_t* func, int fd);

/*
 * Appends a constant to Function.
 *
 * Args:
 *     `value_mode` can only be
 *       - MLVM_COPY_VALUE, which copies the value.
 *       - MLVM_ALIAS_VALUE, which alias the value,
 *       - MLVM_MOVE_VALUE, which moves the states from the original tensor.
 *         After this the original value is invalid for usage
 *
 * Returns:
 *     NULL for error. The returned operand is owned by the function.
 */
extern ir_operand_t*     ir_function_append_constant(ir_function_t* func,
                                                     tensor_t*      tensor,
                                                     int            value_mode);
extern ir_instruction_t* ir_function_append_instruction(
    ir_function_t* func, ir_instruction_type type);

#endif
