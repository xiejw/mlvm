#include "mlvm/ir/function.h"

#include <stdlib.h>

#include "mlvm/ir/error.h"
#include "mlvm/ir/macros.h"

ir_instruction_t* ir_instruction_create(struct ir_function_t* parent_func,
                                        ir_instruction_type   type,
                                        const char*           name_fmt, ...) {
  ir_instruction_t* instruction = malloc(sizeof(ir_instruction_t));
  char*             name;

  MLVM_IR_FILL_NAME(name, name_fmt);

  instruction->name        = name;
  instruction->parent_func = parent_func;
  instruction->type        = type;
  instruction->finalized_  = 0;

  list_init(&instruction->operands);
  list_init(&instruction->outputs);
  return instruction;
}

void ir_instruction_free(ir_instruction_t* instruction) {
  free(instruction->name);
  list_deinit(&instruction->operands);
  list_deinit_with_deleter(&instruction->outputs, ir_operand_free);
  free(instruction);
}

void ir_instruction_append_operand(ir_instruction_t* instruction,
                                   ir_operand_t*     operand) {
  assert(!instruction->finalized_);
  list_append(&instruction->operands, operand);
}

int ir_instruction_finalize(ir_instruction_t* instruction) {
  assert(!instruction->finalized_);
  instruction->finalized_ = 1;

  switch (instruction->type) {
    case IR_OP_ADD:
      return ir_context_set_error(instruction->parent_func->ctx,
                                  MLVM_ERROR_UNSUPPORTED_INSTRUCTION_TYPE,
                                  "Instruction type %d is not supported yet",
                                  instruction->type);
  }

  /* TODO: Add output */
  return 0;
}
