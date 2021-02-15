#ifndef OBJECT_H_
#define OBJECT_H_

#include <inttypes.h>  // int64_t
#include <stdlib.h>    // siz222

#include "adt/sds.h"

enum obj_dtype_t {
        OBJ_DTYPE_SHAPE,
        OBJ_DTYPE_FLOAT32,
};

struct obj_tensor_t {
        int              rank : 6;   // length of dims
        int              owned : 1;  // if 1, own the buffer.
        int              mark : 1;   // gabage collector.
        enum obj_dtype_t dtype : 3;  // data type.
        size_t           size;       // count of elements.
        void *           buffer;     // point to back data. NULL for shape.
        int              dims[];     // dimenions.
};

extern struct obj_tensor_t *objShapeNew(int rank, int dims[]);
extern void                 objShapeFree(struct obj_tensor_t *);

extern struct obj_tensor_t *objTensorNew(enum obj_dtype_t dtype, int rank,
                                         int dims[]);
extern void                 objTensorFree(struct obj_tensor_t *t);

#define objTensorNewFloat32(r, ...) \
        objTensorNew(OBJ_DTYPE_FLOAT32, r, ((int[]){__VA_ARGS__}))

extern void objTensorAllocAndCopy(struct obj_tensor_t *, void *);
extern void objTensorDump(struct obj_tensor_t *, _mut_ sds_t *);

#endif
