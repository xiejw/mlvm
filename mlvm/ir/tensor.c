#include "mlvm/ir/tensor.h"

#include <assert.h>
#include <inttypes.h> /* PRIxN */
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

extern tensor_t* tensor_create(tensor_shape_t rank, tensor_shape_t* shape,
                               double* value, int value_mode) {
  /* Consider to check overflow of the size. */
  tensor_size_t   size;
  tensor_size_t*  stride;
  tensor_shape_t  i;
  tensor_shape_t* shape_copy;
  double*         value_buffer;

  tensor_t* tensor;

  assert(rank >= 1);
  assert(value != NULL);
  assert(value_mode == MLVM_COPY_VALUE || value_mode == MLVM_MOVE_VALUE ||
         value_mode == MLVM_ALIAS_VALUE);

  tensor = malloc(sizeof(tensor_t));

  shape_copy = malloc(rank * sizeof(tensor_shape_t));
  memcpy(shape_copy, shape, rank * sizeof(tensor_shape_t));
  stride = malloc(rank * sizeof(tensor_size_t));

  /* Compuate the size and stride. */
  i    = rank - 1;
  size = 1;
  for (;;) {
    assert(shape[i] > 0);
    stride[i] = size;
    size *= shape[i];
    if (i-- == 0) break;
  }

  if (value_mode == MLVM_COPY_VALUE) {
    tensor_size_t size_of_buffer = size * sizeof(double);
    double*       value_copy     = malloc(size_of_buffer);
    memcpy(value_copy, value, size_of_buffer);
    value_buffer = value_copy;
  } else {
    value_buffer = value;
  }

  tensor->size   = size;
  tensor->rank   = rank;
  tensor->shape  = shape_copy;
  tensor->stride = stride;
  tensor->value  = value_buffer;
  tensor->value_mode_ =
      value_mode == MLVM_ALIAS_VALUE ? MLVM_ALIAS_VALUE : MLVM_OWNING_VALUE;
  return tensor;
}

void tensor_set_stride(tensor_t* tensor, tensor_size_t* new_stride) {
  memcpy(tensor->stride, new_stride, tensor->rank * sizeof(*tensor->stride));
}

void tensor_free(tensor_t* tensor) {
  if (tensor->value_mode_ != MLVM_DEAD_VALUE) {
    free(tensor->shape);
    free(tensor->stride);
    if (tensor->value_mode_ != MLVM_ALIAS_VALUE) free(tensor->value);
  }
  free(tensor);
}

tensor_t* tensor_move(tensor_t* src) {
  tensor_t* dst;
  assert(src->value_mode_ == MLVM_OWNING_VALUE);

  dst = malloc(sizeof(tensor_t));
  memcpy(dst, src, sizeof(tensor_t));

  src->value       = NULL;
  src->value_mode_ = MLVM_DEAD_VALUE;
  return dst;
}

tensor_t* tensor_copy(tensor_t* src) {
  tensor_t* new_tensor;
  new_tensor =
      tensor_create(src->rank, src->shape, src->value, MLVM_COPY_VALUE);
  tensor_set_stride(new_tensor, src->stride);
  return new_tensor;
}

tensor_t* tensor_alias(tensor_t* src) {
  tensor_t* new_tensor;
  new_tensor =
      tensor_create(src->rank, src->shape, src->value, MLVM_ALIAS_VALUE);
  tensor_set_stride(new_tensor, src->stride);
  return new_tensor;
}

int tensor_print(tensor_t* tensor, int fd) {
  int           n = 0;
  tensor_size_t i;
  tensor_size_t size = tensor->size;
  double*       buf  = tensor->value;

  /* Print headline with shape and stride */
  n += dprintf(fd, "Tensor: ");
  n += tensor_print_shape_info(tensor, fd);
  n += dprintf(fd, "\n");

  /* Printf value buffer. */
  n += dprintf(fd, "[ ");
  for (i = 0; i < size; i++) {
    n += dprintf(fd, "%6.3f  ", buf[i]);
    if (i % 10 == 9) n += dprintf(fd, i != size - 1 ? "\n  " : "\n");
  }
  n += dprintf(fd, "]\n");
  return n;
}

int tensor_print_shape_info(tensor_t* tensor, int fd) {
  int            n    = 0;
  tensor_shape_t rank = tensor->rank;
  tensor_shape_t j;

  n += dprintf(fd, "<");
  for (j = 0; j < rank - 1; j++) {
    n += dprintf(fd, "%3d,", tensor->shape[j]);
  }
  n += dprintf(fd, "%3d", tensor->shape[rank - 1]);
  n += dprintf(fd, ">");

  n += dprintf(fd, " /{");
  for (j = 0; j < rank - 1; j++) {
    n += dprintf(fd, "%3" PRIu64 ",", tensor->stride[j]);
  }
  n += dprintf(fd, "%3" PRIu64, tensor->stride[rank - 1]);
  n += dprintf(fd, "}/");
  return n;
}
