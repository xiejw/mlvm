#ifndef MLVM_RUNTIME_KERNEL_SHAPEUTIL_H_
#define MLVM_RUNTIME_KERNEL_SHAPEUTIL_H_

#include "mlvm/ir/tensor.h"

/* Returns 1 if stride is same. */
extern int kernel_stripe_identical(tensor_t* arg_1, tensor_t* arg_2);

#endif
