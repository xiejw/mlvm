#ifndef MLVM_IR_TENSOR_H_
#define MLVM_IR_TENSOR_H_

#include <stdint.h>

#define MLVM_COPY_VALUE 0 /* Copy value into Tensor. */
#define MLVM_MOVE_VALUE 1 /* Move value into Tensor. */

typedef struct {
  uint64_t  size;
  uint64_t  rank;  /* Must be positive (non-zero). */
  uint64_t* shape; /* length is `rank` above. */
  double*   value; /* The value buffer. */
} mlvm_tensor_t;

/* The shape will be copied by the value is transfer to tensor. */
extern mlvm_tensor_t* mlvm_tensor_create(uint64_t rank, uint64_t* shape,
                                         double* value, int value_mode);
extern void           mlvm_tensor_free(mlvm_tensor_t* tensor);

#endif
