#ifndef OBJECT_H_
#define OBJECT_H_

#include <inttypes.h>  // int64_t
#include <stdlib.h>    // sizt_t

#include "adt/sds.h"

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

// TODO(xiejw): remove this.
extern void *obj_tensor_pool;

extern struct obj_tensor_t *objShapeNew(int rank, int dims[]);
extern void                 objShapeFree(struct obj_tensor_t *);

extern struct obj_tensor_t *objTensorNew(int rank, int dims[]);
extern void                 objTensorFree(struct obj_tensor_t *t);

extern void objTensorAllocateAndCopy(struct obj_tensor_t *, obj_float_t *);
extern void objTensorDump(struct obj_tensor_t *, sds_t *);

extern int objGC();

#endif
