#include "mlvm/ir/ir.h"

#include <stdlib.h>

#include "mlvm/ir/macros.h"

static void ir_output_info_free(ir_output_info_t* info) {
  free(info->shape);
  free(info);
}

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

ir_operand_t* ir_operand_create_output(ir_output_info_t* output_info,
                                       const char*       name_fmt, ...) {
  ir_operand_t* operand;
  char*         name;

  MLVM_IR_FILL_NAME(name, name_fmt);

  operand                    = malloc(sizeof(ir_operand_t));
  operand->type              = IR_OPERAND_OUTPUT;
  operand->value.output_info = output_info;
  operand->name              = name;
  return operand;
}

void ir_operand_free(ir_operand_t* operand) {
  switch (operand->type) {
    case IR_OPERAND_CONST:
      tensor_free(operand->value.tensor);
      break;
    case IR_OPERAND_OUTPUT:
      ir_output_info_free(operand->value.output_info);
      break;
  }
  free(operand->name);
  free(operand);
}
