#ifndef OBJECT_H_
#define OBJECT_H_

#include <inttypes.h>  // int64_t
#include <stdlib.h>    // sizt_t

typedef float obj_float_t;

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

typedef struct {
  int rank;
  int dims[];
} obj_shape_t;

typedef struct {
  size_t      size;
  obj_float_t value[];
} obj_array_t;

extern obj_t* objNewInt(int64_t v);
extern obj_t* objNewShape(int rank, int dims[]);
extern obj_t* objNewArray(size_t size, obj_float_t dims[]);

extern void objDecrRefCount(obj_t* o);

#define objInt(o)   (*(int64_t*)(((o) + 1)))
#define objShape(o) (((obj_shape_t*)((o) + 1)))
#define objArray(o) (((obj_array_t*)((o)->ptr)))

#define objIsInt(o)   ((o)->kind == OBJ_INT)
#define objIsShape(o) ((o)->kind == OBJ_SHAPE)
#define objIsArray(o) ((o)->kind == OBJ_ARRAY)

#endif
