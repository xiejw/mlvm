#include "mlvm/ir/tensor.h"

#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

extern tensor_t* tensor_create(uint32_t rank, uint32_t* shape, double* value,
                               int value_mode) {
  /* Consider to check overflow of the size. */
  uint64_t  size;
  uint32_t  i;
  uint32_t* shape_copy;
  double*   value_buffer;

  tensor_t* tensor = malloc(sizeof(tensor_t));
  assert(rank >= 1);

  shape_copy = malloc(rank * sizeof(uint32_t));
  memcpy(shape_copy, shape, rank * sizeof(uint32_t));

  assert(shape[0] > 0);
  size = shape[0];
  for (i = 1; i < rank; i++) {
    assert(shape[i] > 0);
    size *= shape[i];
  }

  if (value_mode == MLVM_COPY_VALUE) {
    uint64_t size_of_buffer = size * sizeof(double);
    double*  value_copy     = malloc(size_of_buffer);
    memcpy(value_copy, value, size_of_buffer);
    value_buffer = value_copy;
  } else {
    value_buffer = value;
  }

  tensor->size        = size;
  tensor->rank        = rank;
  tensor->shape       = shape_copy;
  tensor->value       = value_buffer;
  tensor->value_mode_ = value_mode;
  return tensor;
}

void tensor_free(tensor_t* tensor) {
  free(tensor->shape);
  if (tensor->value_mode_ != MLVM_ALIAS_VALUE) free(tensor->value);
  free(tensor);
}

int tensor_print(tensor_t* tensor, int fd) {
  int      n = 0;
  uint64_t i;
  uint64_t size = tensor->size;
  uint32_t j;
  uint32_t rank = tensor->rank;
  double*  buf  = tensor->value;

  /* Print headline. */
  n += dprintf(fd, "Tensor: <");
  for (j = 0; j < rank - 1; j++) {
    n += dprintf(fd, "%3d,", tensor->shape[j]);
  }
  n += dprintf(fd, "%3d", tensor->shape[rank - 1]);
  n += dprintf(fd, ">\n[ ");

  /* Printf value buffer. */
  for (i = 0; i < size; i++) {
    n += dprintf(fd, "%6.3f  ", buf[i]);
    if (i % 10 == 9) n += dprintf(fd, i != size - 1 ? "\n  " : "\n");
  }
  n += dprintf(fd, "]\n");
  return n;
}
