#ifndef MLVM_IR_TENSOR_H_
#define MLVM_IR_TENSOR_H_

#include <stdint.h>

#define MLVM_COPY_VALUE  0 /* Copy value into Tensor. */
#define MLVM_MOVE_VALUE  1 /* Move value into Tensor. */
#define MLVM_ALIAS_VALUE 2 /* Alias value, which must have longer life time. \
                            */

typedef struct {
  uint64_t  size;  /* Total number of elements. */
  uint32_t  rank;  /* Must be positive (non-zero). */
  uint32_t* shape; /* Length is `rank` above. */
  double*   value; /* The value buffer. */

  /* Internal fields */
  int value_mode_;
} tensor_t;

extern tensor_t* tensor_create(uint32_t rank, uint32_t* shape, double* value,
                               int value_mode);
extern void      tensor_free(tensor_t* tensor);

extern int tensor_print(tensor_t* tensor, int fd);

#endif
