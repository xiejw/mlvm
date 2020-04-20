#ifndef MLVM_IR_TENSOR_H_
#define MLVM_IR_TENSOR_H_

#include <stdint.h>

#define MLVM_COPY_VALUE  0 /* Copy value into Tensor. */
#define MLVM_MOVE_VALUE  1 /* Move value into Tensor. */
#define MLVM_ALIAS_VALUE 2 /* Alias value, who must have longer life time. */

typedef struct {
  uint64_t  size;
  uint32_t  rank;  /* Must be positive (non-zero). */
  uint32_t* shape; /* length is `rank` above. */
  double*   value; /* The value buffer. */

  /* Internal fields */
  int value_mode_;
} mlvm_tensor_t;

/* The shape will be copied by the value is transfer to tensor. */
extern mlvm_tensor_t* mlvm_tensor_create(uint32_t rank, uint32_t* shape,
                                         double* value, int value_mode);
extern void           mlvm_tensor_free(mlvm_tensor_t* tensor);

extern int mlvm_tensor_print(mlvm_tensor_t* tensor, int fd);

#endif