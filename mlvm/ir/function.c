#include "mlvm/ir/function.h"

#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

ir_function_t* ir_function_create(ir_context_t* ctx, char* name) {
  ir_function_t* func      = malloc(sizeof(ir_function_t));
  size_t         name_size = strlen(name);
  func->name               = malloc((name_size + 1) * sizeof(char));
  func->ctx                = ctx;

  /* Init name. */
  strcpy(func->name, name);

  /* Initialize all lists. */
  list_init(&func->const_tensors);
  list_init(&func->instructions);

  return func;
}

void ir_function_free(ir_function_t* func) {
  free(func->name);
  list_deinit_with_deleter(&func->const_tensors, ir_operand_free);
  list_deinit_with_deleter(&func->instructions, ir_instruction_free);
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

  { /* Prints out the instructions. */
    list_ir_instruction_t* instructions = &func->instructions;
    if (list_size(instructions) > 0) {
      mlvm_size_t i, size = list_size(instructions);
      n += dprintf(fd, "  Instructions:\n");

      for (i = 0; i < size; i++) {
        ir_instruction_t* instruction = instructions->data[i];
        n += dprintf(fd, "    Name `%s`: ", instruction->name);
        n += dprintf(fd, "\n");
      }
    }
  }

  return n;
}

ir_operand_t* ir_function_append_constant(ir_function_t* func, tensor_t* tensor,
                                          int value_mode) {
  ir_operand_t* operand =
      ir_operand_create_const(tensor_clone(tensor, value_mode), "const_%d",
                              (int)list_size(&func->const_tensors));

  list_append(&func->const_tensors, operand);
  return operand;
}

ir_instruction_t* ir_function_append_instruction(ir_function_t*      func,
                                                 ir_instruction_type type) {
  ir_instruction_t* instruction = ir_instruction_create(
      func, type, "inst_%d", (int)list_size(&func->instructions));
  list_append(&func->instructions, instruction);
  return instruction;
}
