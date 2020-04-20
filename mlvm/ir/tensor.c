#include "mlvm/ir/tensor.h"

#include <assert.h>
#include <stdlib.h>
#include <string.h>

extern mlvm_tensor_t* mlvm_tensor_create(uint64_t rank, uint64_t* shape,
                                         double* value, int value_mode) {
  /* Consider to check overflow of the size. */
  uint64_t  size;
  uint64_t  i;
  uint64_t* shape_copy;
  double*   value_buffer;

  mlvm_tensor_t* tensor = malloc(sizeof(mlvm_tensor_t));
  assert(rank >= 1);

  shape_copy = malloc(rank * sizeof(uint64_t));
  memcpy(shape_copy, shape, rank * sizeof(uint64_t));

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

  tensor->size  = size;
  tensor->rank  = rank;
  tensor->shape = shape_copy;
  tensor->value = value_buffer;
  return NULL;
}

void mlvm_tensor_free(mlvm_tensor_t* tensor) {
  free(tensor->shape);
  free(tensor->value);
  free(tensor);
}
