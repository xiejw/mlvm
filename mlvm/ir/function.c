#include "mlvm/ir/function.h"

#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_TENSOR_NAME 128

static void ir_function_free_operands(list_ir_operand_t* operands) {
  mlvm_size_t size = list_size(operands);
  mlvm_size_t i;
  for (i = 0; i < size; i++) {
    ir_operand_t* operand = operands->data[i];
    if (operand->type == IR_CONST) tensor_free(operand->value.tensor);
    free(operand->name);
    free(operand);
  }
  list_deinit(operands);
}

ir_function_t* ir_function_create(char* name) {
  ir_function_t* func      = malloc(sizeof(ir_function_t));
  size_t         name_size = strlen(name);
  func->name               = malloc((name_size + 1) * sizeof(char));

  /* Init name. */
  strcpy(func->name, name);
  /* Init const tensors. */
  list_init(&func->const_tensors);

  return func;
}

void ir_function_free(ir_function_t* func) {
  free(func->name);
  ir_function_free_operands(&func->const_tensors);
  free(func);
}

int ir_function_print(ir_function_t* func, int fd) {
  int n = 0;
  n += dprintf(fd, "Func: %s\n", func->name);

  { /* Prints out the contants. */
    list_ir_operand_t* const_tensors = &func->const_tensors;
    if (list_size(const_tensors) > 0) {
      mlvm_size_t i, size = list_size(const_tensors);
      n += dprintf(fd, "  Constants:\n");

      for (i = 0; i < size; i++) {
        ir_operand_t* operand = const_tensors->data[i];
        tensor_t*     tensor  = operand->value.tensor;
        n += dprintf(fd, "    Name `%s`: ", operand->name);
        n += tensor_print_shape_info(tensor, fd);
        n += dprintf(fd, "\n");
      }
    }
  }

  return n;
}

ir_operand_t* ir_function_append_constant(ir_function_t* func, tensor_t* tensor,
                                          int value_mode) {
  tensor_t*     const_tensor;
  ir_operand_t* operand;

  assert(value_mode == MLVM_COPY_VALUE || value_mode == MLVM_MOVE_VALUE ||
         value_mode == MLVM_ALIAS_VALUE);

  switch (value_mode) {
    case MLVM_COPY_VALUE:
      const_tensor = tensor_copy(tensor);
      break;
    case MLVM_MOVE_VALUE:
      const_tensor = tensor_move(tensor);
      break;
    case MLVM_ALIAS_VALUE:
      const_tensor = tensor_alias(tensor);
      break;
  }

  operand               = malloc(sizeof(ir_operand_t));
  operand->type         = IR_CONST;
  operand->value.tensor = const_tensor;

  /* Fill the name. */
  {
    int   size = (int)list_size(&func->const_tensors);
    char* name = malloc(MAX_TENSOR_NAME * sizeof(char));
    sprintf(name, "const_%d", size);
    operand->name = name;
  }

  list_append(&func->const_tensors, operand);

  return operand;
}
