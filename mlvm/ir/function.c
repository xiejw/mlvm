#include "mlvm/ir/function.h"

#include <assert.h>
#include <stdlib.h>
#include <string.h>

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

  {
    list_tensor_t* const_tensors = &func->const_tensors;
    uint64_t       size          = list_size(const_tensors);
    uint64_t       i;
    for (i = 0; i < size; i++) {
      tensor_free(const_tensors->data[i]);
    }
    list_deinit(const_tensors);
  }

  free(func);
}

ir_operand_t* ir_function_add_constant(ir_function_t* func, tensor_t* tensor,
                                       int value_mode) {
  assert(value_mode == MLVM_COPY_VALUE || value_mode == MLVM_MOVE_VALUE ||
         value_mode == MLVM_ALIAS_VALUE);

  (void)func;
  (void)tensor;
  return NULL;
}
