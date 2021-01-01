#include "object.h"

#include <assert.h>
#include <stdlib.h>
#include <string.h>

#include "adt/vec.h"
#include "base/error.h"

// embeded int.
obj_t* objNewInt(int64_t v) {
  obj_t*   o   = malloc(sizeof(obj_t) + sizeof(int64_t));
  int64_t* buf = (int64_t*)(o + 1);

  o->kind      = OBJ_INT;
  o->ref_count = 1;
  o->ptr       = buf;
  *buf         = v;
  return o;
}

// embeded shape.
obj_t* objNewShape(int rank, int dims[]) {
  assert(rank > 0);
  obj_t* o = malloc(sizeof(obj_t) + sizeof(obj_shape_t) + sizeof(int) * rank);
  obj_shape_t* buf = (obj_shape_t*)(o + 1);

  o->kind      = OBJ_SHAPE;
  o->ref_count = 1;
  o->ptr       = buf;

  buf->rank = rank;
  memcpy(buf->dims, dims, rank * sizeof(int));
  return o;
}

// embeded array.
static inline obj_t* objNewEmbeddingArray(size_t size, obj_float_t value[]) {
  assert(size > 0);
  obj_t* o =
      malloc(sizeof(obj_t) + sizeof(obj_array_t) + sizeof(obj_float_t) * size);
  obj_array_t* buf = (obj_array_t*)(o + 1);

  o->kind      = OBJ_ARRAY;
  o->ref_count = 1;
  o->ptr       = buf;

  buf->size = size;
  memcpy(buf->value, value, size * sizeof(obj_float_t));
  return o;
}

obj_t* objNewArray(size_t size, obj_float_t value[]) {
  return objNewEmbeddingArray(size, value);
}

void objDecrRefCount(obj_t* o) {
  if (o == NULL) return;

  if (!--(o->ref_count)) {
    switch (o->kind) {
      case OBJ_INT:
      case OBJ_SHAPE:
      case OBJ_ARRAY:
        // TODO non-embedding array
        free(o);
        break;
      default:
        errFatalAndExit("objDecrRefCount unknown object kind: %d\n", o->kind);
    }
  }
}
