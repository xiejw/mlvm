#include <stdio.h>

#include "mlvm/ir/ir.h"
#include "mlvm/sprng/sprng.h"

tensor_t* create_a_random_tensor(sprng_t* prng) {
  mlvm_uint_t rank    = 2;
  uint32_t    shape[] = {10, 10};
  tensor_t*   tensor;
  double*     value;

  /* This value moved into the tensor. */
  value = malloc(100 * sizeof(double));
  srng_standard_normal(prng, 100, value);
  tensor = tensor_create(rank, shape, value, MLVM_MOVE_VALUE);
  return tensor;
}

int build_simple_func(ir_function_t* func, sprng_t* prng) {
  int               err;
  tensor_t*         tensor;
  ir_instruction_t* ins;
  ir_operand_t*     operand;

  tensor  = create_a_random_tensor(prng);
  operand = ir_function_append_constant(func, tensor, MLVM_MOVE_VALUE);
  tensor_free(tensor);
  if (operand == NULL) return -1;

  ins = ir_function_append_instruction(func, IR_OP_ADD);
  ir_instruction_append_operand(ins, operand);
  ir_instruction_append_operand(ins, operand);

  err = ir_instruction_finalize(ins);
  if (err) return err;

  ir_function_print(func, 1);
  return 0;
}

int main() {
  int            err;
  ir_context_t*  ctx = ir_context_create();
  ir_function_t* func;

  ir_context_set_prng(ctx, sprng_create(456L));

  func = ir_function_create(ctx, "main");

  if ((err = build_simple_func(func, ctx->prng)))
    fprintf(stderr, "Unexpected error: %d: %s\n", err, ctx->error_message);

  ir_function_free(func);
  ir_context_free(ctx);
  return 0;
}
