#include "object.h"

#include <assert.h>
#include <stdlib.h>
#include <string.h>

#include "adt/vec.h"
#include "base/error.h"

// embeded int.
obj_t* objNewInt64(int64_t v) {
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

  o->kind      = OBJ_INT;
  o->ref_count = 1;
  o->ptr       = buf;

  buf->rank = rank;
  memcpy(buf->dims, dims, rank * sizeof(int));
  return o;
}

void objDecrRefCount(obj_t* o) {
  if (o == NULL) return;

  if (!--(o->ref_count)) {
    switch (o->kind) {
      case OBJ_INT:
        free(o);
        break;
      default:
        errFatalAndExit("unknown object kind: %d", o->kind);
    }
  }
}
