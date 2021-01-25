#include "object.h"

#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "adt/vec.h"
#include "base/error.h"

// -----------------------------------------------------------------------------
// internal data structure.
// -----------------------------------------------------------------------------

// used to detect whether a tensor is a shape.
static obj_float_t shape_indicator[1];

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
        o->size   = eleEize(rank, dims);
        o->buffer = shape_indicator;
        memcpy(o->dims, dims, rank * sizeof(int));
        return o;
}

void objShapeFree(struct obj_tensor_t* t)
{
        if (t == NULL) return;
        assert(!t->owned);
        assert(t->buffer == shape_indicator);
        free(t);
}

struct obj_tensor_t* objTensorNew(int rank, int dims[])
{
        struct obj_tensor_t* o =
            malloc(sizeof(struct obj_tensor_t) + rank * sizeof(int));
        o->rank   = rank;
        o->owned  = 0;
        o->mark   = 0;
        o->buffer = NULL;
        o->size   = eleEize(rank, dims);
        memcpy(o->dims, dims, rank * sizeof(int));

        return o;
}

void objTensorFree(struct obj_tensor_t* t)
{
        if (t == NULL) return;
        if (t->owned) free(t->buffer);
        free(t);
}

// If buf is NULL, copy will not happen.
void objTensorAllocAndCopy(struct obj_tensor_t* t, obj_float_t* buf)
{
        assert(t->owned != 1);
        assert(t->buffer == NULL);
        size_t size = t->size * sizeof(obj_float_t);
        t->owned    = 1;
        t->buffer   = malloc(size);
        if (buf != NULL) memcpy(t->buffer, buf, size);
}

void objTensorDump(struct obj_tensor_t* t, sds_t* s)
{
        assert(t->buffer != shape_indicator);  // use helper method.
        assert(t->buffer != NULL);
        sdsCatPrintf(s, "[ ");
        int size_to_print = t->size;
        if (size_to_print > 10) size_to_print = 10;
        for (int i = 0; i < size_to_print; i++) {
                sdsCatPrintf(s, " %f,", t->buffer[i]);
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
