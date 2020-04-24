#ifndef MLVM_RUNTIME_KERNEL_H_
#define MLVM_RUNTIME_KERNEL_H_

#include "mlvm/ir/tensor.h"

/* Kernels do not indicate errors. Callers must obey the shape/stride rules. */

/* add.c */
extern void kernel_add(tensor_t* output, tensor_t* arg_1, tensor_t* arg_2);
/* mul.c */
extern void kernel_mul(tensor_t* output, tensor_t* arg_1, tensor_t* arg_2);
/* mamul.c */
extern void kernel_matmul(tensor_t* output, tensor_t* arg_1, tensor_t* arg_2);

#endif
