#ifndef MLVM_IR_TENSOR_H_
#define MLVM_IR_TENSOR_H_

#include <stdint.h>

#include "mlvm/lib/types.h"

#define MLVM_COPY_VALUE   0 /* Copy value into Tensor. */
#define MLVM_MOVE_VALUE   1 /* Move value into Tensor. */
#define MLVM_ALIAS_VALUE  2 /* Alias, which must have longer life time. */
#define MLVM_OWNING_VALUE 3 /* Owning the value. */
#define MLVM_DEAD_VALUE   4 /* Dangling value. */

typedef struct {
  mlvm_uint_t  rank;   /* Must be positive (non-zero). */
  mlvm_uint_t* shape;  /* Length is `rank` above. */
  mlvm_size_t* stride; /* Stride for each dim. */
  mlvm_size_t  size;   /* Total number of elements. */
  double*      value;  /* The value buffer. */

  /* Internal fields */
  int value_mode_; /* Can only be MLVM_ALIAS_VALUE, MLVM_OWNING_VALUE,
                      MLVM_DEAD_VALUE */
} tensor_t;

/* value_mode can only be MLVM_COPY_VALUE, MLVM_MOVE_VALUE, or MLVM_ALIAS_VALUE.
 */
extern tensor_t* tensor_create(mlvm_uint_t rank, mlvm_uint_t* shape,
                               double* value, int value_mode);
extern void      tensor_free(tensor_t* tensor);
extern int       tensor_print(tensor_t* tensor, int fd);
extern int       tensor_print_shape_info(tensor_t* tensor, int fd);

/* Copy the new stride into the tensor struct. */
extern void tensor_set_stride(tensor_t* tensor, mlvm_size_t* new_stride);

/*
 * Returns a _copy_ of tensor `src`.
 *
 * All fields are copied, (see the exception of MLVM_MOVE_VALUE) while the
 * value field is controlled by `value_mode`.
 *
 * Args:
 *    value_mode can only be MLVM_COPY_VALUE, MLVM_MOVE_VALUE, or
 *    MLVM_ALIAS_VALUE. In particular,
 *      - MLVM_COPY_VALUE will copy the value (which has runtime cost).
 *      - MLVM_ALIAS_VALUE will alias the value. It is caller's responsibility
 *        to ensure the lifttime of the src is longer than the copied return
 *        value.
 *      - MLVM_MOVE_VALUE will steal all fields from the `src` given the `src`
 *        will be left as a dead state. In addition, the
 *        1. The `src` must own its value, i.e., value_mode_ should be
 *           MLVM_OWNING_VALUE.
 *        2. `tensor_free` must be called to free the src itself.
 *
 */
extern tensor_t* tensor_clone(tensor_t* src, int value_mode);

#endif
