#ifndef MLVM_IR_TENSOR_H_
#define MLVM_IR_TENSOR_H_

#include <stdint.h>

typedef uint32_t tensor_shape_t;
typedef uint64_t tensor_size_t;

#define MLVM_COPY_VALUE  0 /* Copy value into Tensor. */
#define MLVM_MOVE_VALUE  1 /* Move value into Tensor. */
#define MLVM_ALIAS_VALUE 2 /* Alias, which must have longer life time. */

typedef struct {
  tensor_shape_t  rank;   /* Must be positive (non-zero). */
  tensor_shape_t* shape;  /* Length is `rank` above. */
  tensor_size_t*  stride; /* Stride for each dim. */
  tensor_size_t   size;   /* Total number of elements. */
  double*         value;  /* The value buffer. */

  /* Internal fields */
  int value_mode_;
} tensor_t;

extern tensor_t* tensor_create(tensor_shape_t rank, tensor_shape_t* shape,
                               double* value, int value_mode);
/* Copy the new stride into the tensor struct. */
extern void tensor_set_stride(tensor_t* tensor, tensor_size_t* new_stride);
extern void tensor_free(tensor_t* tensor);

extern int tensor_print(tensor_t* tensor, int fd);

#endif
