#include "mlvm/ir/ir.h"

#include <stdlib.h>

#include "mlvm/ir/macros.h"

ir_operand_t* ir_operand_create_const(tensor_t*   const_tensor,
                                      const char* name_fmt, ...) {
  ir_operand_t* operand;
  char*         name;

  MLVM_IR_FILL_NAME(name, name_fmt);

  operand               = malloc(sizeof(ir_operand_t));
  operand->type         = IR_OPERAND_CONST;
  operand->value.tensor = const_tensor;
  operand->name         = name;
  return operand;
}

void ir_operand_free(ir_operand_t* operand) {
  if (operand->type == IR_OPERAND_CONST) tensor_free(operand->value.tensor);
  free(operand->name);
  free(operand);
}
