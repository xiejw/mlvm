#ifndef OBJECT_H_
#define OBJECT_H_

#include <inttypes.h>  // int64_t

typedef enum {
  OBJ_INT,
  OBJ_SHAPE,
  OBJ_ARRAY,
  OBJ_TENSOR,
} obj_kind_t;

typedef struct {
  obj_kind_t kind;
  int        ref_count;
  void*      ptr;
} obj_t;

extern obj_t* objNewInt64(int64_t v);
extern void   objDecrRefCount(obj_t* o);

#define objInt64Value(o) (*(int64_t*)(((o) + 1)))

#endif
