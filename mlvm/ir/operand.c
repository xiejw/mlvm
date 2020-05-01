#include "mlvm/ir/ir.h"

#include <stdarg.h>
#include <stdlib.h>

#define MAX_TENSOR_NAME 128

ir_operand_t* ir_operand_create_const(tensor_t*   const_tensor,
                                      const char* name_fmt, ...) {
  ir_operand_t* operand;
  char*         name;

  {
    va_list args;
    va_start(args, name_fmt);
    name = malloc(MAX_TENSOR_NAME * sizeof(char));
    sprintf(name, name_fmt, args);
    va_end(args);
  }

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
