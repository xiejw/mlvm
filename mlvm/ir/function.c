#include "mlvm/ir/function.h"

#include <stdlib.h>
#include <string.h>

ir_function_t* ir_function_create(char* name) {
  ir_function_t* func      = malloc(sizeof(ir_function_t));
  size_t         name_size = strlen(name);
  func->name               = malloc((name_size + 1) * sizeof(char));

  strcpy(func->name, name);

  return func;
}

void ir_function_free(ir_function_t* func) {
  free(func->name);
  free(func);
}
