#include "mlvm/ir/ir.h"

#include <stdlib.h>

/*
static ir_operand_t* ir_operand_create_const() {
  operand               = malloc(sizeof(ir_operand_t));
  operand->type         = IR_CONST;
  operand->value.tensor = const_tensor;

  {
    int   size = (int)list_size(&func->const_tensors);
    char* name = malloc(MAX_TENSOR_NAME * sizeof(char));
    sprintf(name, "const_%d", size);
    operand->name = name;
  }
}
*/

void ir_operand_free(ir_operand_t* operand) {
  if (operand->type == IR_OPERAND_CONST) tensor_free(operand->value.tensor);
  free(operand->name);
  free(operand);
}
