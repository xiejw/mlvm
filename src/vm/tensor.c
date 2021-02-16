#include "tensor.h"

#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "adt/vec.h"
#include "base/error.h"

// -----------------------------------------------------------------------------
// internal data structure.
// -----------------------------------------------------------------------------

// unused.

// -----------------------------------------------------------------------------
// internal prototype.
// -----------------------------------------------------------------------------
size_t eleEize(int rank, int dims[]);

// -----------------------------------------------------------------------------
// implementation.
// -----------------------------------------------------------------------------

// consider to do some optimization to do lookup.
struct obj_tensor_t* objShapeNew(int rank, int dims[])
{
        struct obj_tensor_t* o =
            malloc(sizeof(struct obj_tensor_t) + rank * sizeof(int));
        o->rank   = rank;
        o->owned  = 0;
        o->mark   = 0;
        o->dtype  = OBJ_DTYPE_SHAPE;
        o->size   = eleEize(rank, dims);
        o->buffer = NULL;
        memcpy(o->dims, dims, rank * sizeof(int));
        return o;
}

void objShapeFree(struct obj_tensor_t* t)
{
        if (t == NULL) return;
        assert(!t->owned);
        assert(t->buffer == NULL);
        assert(t->dtype == OBJ_DTYPE_SHAPE);
        free(t);
}

struct obj_tensor_t* objTensorNew(enum obj_dtype_t dtype, int rank, int dims[])
{
        struct obj_tensor_t* o =
            malloc(sizeof(struct obj_tensor_t) + rank * sizeof(int));
        o->rank   = rank;
        o->owned  = 0;
        o->mark   = 0;
        o->dtype  = dtype;
        o->size   = eleEize(rank, dims);
        o->buffer = NULL;
        memcpy(o->dims, dims, rank * sizeof(int));
        return o;
}

void objTensorFree(struct obj_tensor_t* t)
{
        if (t == NULL) return;
        if (t->owned) {
                assert(t->buffer != NULL);
                free(t->buffer);
        } else {
                assert(t->dtype != OBJ_DTYPE_SHAPE);
        }
        free(t);
}

// If buf is NULL, copy will not happen.
void objTensorAllocAndCopy(struct obj_tensor_t* t, void* buf)
{
        assert(t->owned != 1);
        assert(t->buffer == NULL);
        size_t size = 0;

        switch (t->dtype) {
        case OBJ_DTYPE_FLOAT32:
                size = t->size * sizeof(float);
                break;
        case OBJ_DTYPE_INT32:
                size = t->size * sizeof(int32_t);
                break;
        default:
                errFatalAndExit("dtype %d is not supported yet.", t->dtype);
        }

        t->owned  = 1;
        t->buffer = malloc(size);
        if (buf != NULL) memcpy(t->buffer, buf, size);
}

void objTensorDump(struct obj_tensor_t* t, sds_t* s)
{
        if (t->buffer == NULL) {  // fast path
                sdsCatPrintf(s, "[ (NULL) ]");
                return;
        }

        sdsCatPrintf(s, "[");
        int size_to_print = t->size;
        if (size_to_print > 10) size_to_print = 10;

        switch (t->dtype) {
        case OBJ_DTYPE_FLOAT32: {
                float* buf = t->buffer;
                for (int i = 0; i < size_to_print; i++) {
                        sdsCatPrintf(s, " %f,", buf[i]);
                }
        } break;
        case OBJ_DTYPE_INT32: {
                int32_t* buf = t->buffer;
                for (int i = 0; i < size_to_print; i++) {
                        sdsCatPrintf(s, " %d,", buf[i]);
                }
        } break;
        default:
                errFatalAndExit("dtype %d is not supported yet.", t->dtype);
        }

        sdsCatPrintf(s, "]");
}

// -----------------------------------------------------------------------------
// helper method implementation.
// -----------------------------------------------------------------------------

size_t eleEize(int rank, int dims[])
{
        size_t count = 1;
        for (int i = 0; i < rank; i++) {
                count *= dims[i];
        }
        return count;
}
