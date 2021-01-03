#ifndef OBJECT_H_
#define OBJECT_H_

#include <inttypes.h>  // int64_t
#include <stdlib.h>    // sizt_t

typedef float obj_float_t;

typedef enum {
  OBJ_INT,
  OBJ_SHAPE,
  OBJ_TENSOR,
} obj_kind_t;

typedef struct {
  int          rank : 6;   // length of dims
  int          owned : 1;  // if 1, own the buffer.
  int          mark : 1;   // gabage collector.
  obj_float_t *buffer;     // NULL for OBJ_SHAPE.
  int          dims[];     // size of rank.
} obj_tensor_t;

typedef union {
  int64_t       i;
  obj_tensor_t *tensor;
} obj_value_t;

typedef struct {
  obj_kind_t  kind : 7;
  int         marked : 1;  // gabage collector.
  obj_value_t value;
} obj_t;

extern void *obj_tensor_pool;

extern obj_tensor_t *objTensorNew(int rank, int dims[]);
extern void          objTensorFree(obj_tensor_t *t);
extern int           objTensorGabageCollector();

// extern obj_t* objNewInt(int64_t v);
// extern obj_t* objNewShape(int rank, int dims[]);
// extern obj_t* objNewArray(size_t size, obj_float_t dims[]);
//
// extern void objDecrRefCount(obj_t* o);
//
// #define objInt(o)   (*(int64_t*)(((o) + 1)))
// #define objShape(o) (((obj_shape_t*)((o) + 1)))
// #define objArray(o) (((obj_array_t*)((o)->ptr)))
//
// #define objIsInt(o)   ((o)->kind == OBJ_INT)
// #define objIsShape(o) ((o)->kind == OBJ_SHAPE)
// #define objIsArray(o) ((o)->kind == OBJ_ARRAY)

#endif
