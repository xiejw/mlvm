#ifndef OBJECT_H_
#define OBJECT_H_

#include <inttypes.h>  // int64_t
#include <stdlib.h>    // sizt_t

typedef float obj_float_t;

enum obj_kind_t {
        OBJ_INT,
        OBJ_FLOAT,
        OBJ_SHAPE,
        OBJ_TENSOR,
};

struct obj_tensor_t {
        int          rank : 6;   // length of dims
        int          owned : 1;  // if 1, own the buffer.
        int          mark : 1;   // gabage collector.
        size_t       size;       // count of elements.
        obj_float_t *buffer;     // point to internal buf for OBJ_SHAPE.
        int          dims[];     // dimenions.
};

union obj_value_t {
        int64_t              i;
        obj_float_t          f;
        struct obj_tensor_t *t;
};

struct obj_t {
        enum obj_kind_t   kind;
        union obj_value_t value;
};

extern void *obj_tensor_pool;

extern struct obj_tensor_t *objTensorNew(int rank, int dims[]);
extern struct obj_tensor_t *objShapeNew(int rank, int dims[]);
extern void                 objTensorFree(struct obj_tensor_t *t);
extern int                  objGC();

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
