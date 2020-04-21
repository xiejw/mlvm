#ifndef MLVM_RUNTIME_KERNEL_H_
#define MLVM_RUNTIME_KERNEL_H_

#include "mlvm/ir/tensor.h"

/* add.c */
int kernel_add(tensor_t* output, tensor_t* arg_1, tensor_t* arg_2);
/* mul.c */
int kernel_mul(tensor_t* output, tensor_t* arg_1, tensor_t* arg_2);
/* mamul.c */
int kernel_matmul(tensor_t* output, tensor_t* arg_1, tensor_t* arg_2);

#endif
