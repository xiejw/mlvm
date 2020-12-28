#include "object.h"

#include <stdlib.h>

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
