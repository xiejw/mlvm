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

typedef struct obj_tensor_item_t {
        struct obj_tensor_t*      item;
        struct obj_tensor_item_t* next;
} obj_tensor_item_t;

void* obj_tensor_pool = NULL;

// used to detect whether a tensor is a shape.
static obj_float_t shape_indicator[1];

// -----------------------------------------------------------------------------
// implementation.
// -----------------------------------------------------------------------------

int objGC()
{
        if (obj_tensor_pool == NULL) return 0;

        int                count = 0;
        obj_tensor_item_t* p     = obj_tensor_pool;
        obj_tensor_item_t* prev  = NULL;
        while (p != NULL) {
                struct obj_tensor_t* item = p->item;
                if (item->mark) {
                        item->mark = 0;
                        prev       = p;
                        p          = p->next;
                } else {
                        if (prev == NULL) {
                                obj_tensor_pool = p->next;
                        } else {
                                prev->next = p->next;
                        }
                        objTensorFree(item);

                        obj_tensor_item_t* old_p;
                        old_p = p;
                        p     = p->next;
                        free(old_p);
                        count++;
                }
        }
        return count;
}

// consider to do some optimization to do lookup.
struct obj_tensor_t* objShapeNew(int rank, int dims[])
{
        struct obj_tensor_t* o =
            malloc(sizeof(struct obj_tensor_t) + rank * sizeof(int));
        o->rank   = rank;
        o->owned  = 0;
        o->mark   = 0;
        o->buffer = shape_indicator;
        memcpy(o->dims, dims, rank * sizeof(int));

        obj_tensor_item_t* p = malloc(sizeof(obj_tensor_item_t));
        p->item              = o;
        p->next              = obj_tensor_pool;
        obj_tensor_pool      = p;

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
        memcpy(o->dims, dims, rank * sizeof(int));

        obj_tensor_item_t* p = malloc(sizeof(obj_tensor_item_t));
        p->item              = o;
        p->next              = obj_tensor_pool;
        obj_tensor_pool      = p;

        return o;
}

void objTensorFree(struct obj_tensor_t* t)
{
        if (t == NULL) return;
        if (t->owned) free(t->buffer);
        free(t);
}

// #define OBJ_EMBEDDING_ARRAY_SIZE 16
//
// // embeded int.
// obj_t* objNewInt(int64_t v) {
//   obj_t*   o   = malloc(sizeof(obj_t) + sizeof(int64_t));
//   int64_t* buf = (int64_t*)(o + 1);
//
//   o->kind      = OBJ_INT;
//   o->ref_count = 1;
//   o->ptr       = buf;
//   *buf         = v;
//   return o;
// }
//
// // embeded shape.
// obj_t* objNewShape(int rank, int dims[]) {
//   assert(rank > 0);
//   obj_t* o = malloc(sizeof(obj_t) + sizeof(obj_shape_t) + sizeof(int) *
//   rank); obj_shape_t* buf = (obj_shape_t*)(o + 1);
//
//   o->kind      = OBJ_SHAPE;
//   o->ref_count = 1;
//   o->ptr       = buf;
//
//   buf->rank = rank;
//   memcpy(buf->dims, dims, rank * sizeof(int));
//   return o;
// }
//
// // embeded array.
// static inline obj_t* objNewEmbeddingArray(size_t size, obj_float_t value[]) {
//   assert(size > 0);
//   obj_t* o =
//       malloc(sizeof(obj_t) + sizeof(obj_array_t) + sizeof(obj_float_t) *
//       size);
//   obj_array_t* buf = (obj_array_t*)(o + 1);
//
//   o->kind      = OBJ_ARRAY;
//   o->ref_count = 1;
//   o->ptr       = buf;
//
//   buf->size = size;
//   memcpy(buf->value, value, size * sizeof(obj_float_t));
//   return o;
// }
//
// obj_t* objNewArray(size_t size, obj_float_t value[]) {
//   assert(size <= OBJ_EMBEDDING_ARRAY_SIZE);
//   return objNewEmbeddingArray(size, value);
// }
//
// void objDecrRefCount(obj_t* o) {
//   if (o == NULL) return;
//
//   if (!--(o->ref_count)) {
//     switch (o->kind) {
//       case OBJ_INT:
//       case OBJ_SHAPE:
//       case OBJ_ARRAY:
//         // TODO non-embedding array
//         free(o);
//         break;
//       default:
//         errFatalAndExit("objDecrRefCount unknown object kind: %d\n",
//         o->kind);
//     }
//   }
// }
