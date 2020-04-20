#include "mlvm/ir/tensor.h"

#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

extern mlvm_tensor_t* mlvm_tensor_create(uint32_t rank, uint32_t* shape,
                                         double* value, int value_mode) {
  /* Consider to check overflow of the size. */
  uint64_t  size;
  uint32_t  i;
  uint32_t* shape_copy;
  double*   value_buffer;

  mlvm_tensor_t* tensor = malloc(sizeof(mlvm_tensor_t));
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

void mlvm_tensor_free(mlvm_tensor_t* tensor) {
  free(tensor->shape);
  if (tensor->value_mode_ != MLVM_ALIAS_VALUE) free(tensor->value);
  free(tensor);
}

int mlvm_tensor_print(mlvm_tensor_t* tensor, int fd) {
  int      n = 0;
  uint64_t i;
  uint64_t size = tensor->size;
  double*  buf  = tensor->value;

  for (i = 0; i < size; i++) {
    n += dprintf(fd, "%6.3f  ", buf[i]);
    if (i % 10 == 9) n += dprintf(fd, "\n");
  }
  return n;
}
