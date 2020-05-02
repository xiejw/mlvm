#include "mlvm/ir/function.h"

#include <stdlib.h>

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
  return instruction;
}

void ir_instruction_free(ir_instruction_t* instruction) {
  free(instruction->name);
  free(instruction);
}
