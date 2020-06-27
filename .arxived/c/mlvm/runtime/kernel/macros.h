#ifndef MLVM_RUNTIME_KERNEL_MACROS_H_
#define MLVM_RUNTIME_KERNEL_MACROS_H_

#include "mlvm/ir/tensor.h"

#define MLVM_KERNEL_ELEMENT_OP_PLAIN_LOOP(output, arg_1, arg_2, op, size) \
  do {                                                                    \
    mlvm_size_t index;                                                    \
    for (index = 0; index < size; index++) {                              \
      output->value[index] = arg_1->value[index] op arg_2->value[index];  \
    };                                                                    \
  } while (0)

#endif
