#ifndef MLVM_IR_FUNCTION_H_
#define MLVM_IR_FUNCTION_H_

#include <stdint.h>

#include "mlvm/ir/tensor.h"

typedef enum { IR_CONST } ir_operand_type;

/* Does not own any of them. */
typedef union {
  tensor_t* const_tensor;
} ir_operand_value;

typedef struct {
  ir_operand_type  type;
  ir_operand_value value;
} ir_operand_t;

typedef struct {
  char* name; /* Function name. */

} ir_function_t;

ir_function_t* ir_function_create(char* name);
void           ir_function_free(ir_function_t* func);

#endif
