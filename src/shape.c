#include "vm.h"

#include <stdlib.h>

struct shape_t* spNew(int rank, int* dims)
{
        // some thoughts: if rank <= 3, we could build a bst to lookup existing
        // shape and reusing them. This should be fairly cheap to do and save
        // heap allocations.

        if (rank == 0) return NULL;

        struct shape_t* s = malloc(sizeof(struct shape_t) + rank * sizeof(int));
        uint64_t        size = 1;
        for (int i = 0; i < rank; i++) {
                int d = dims[i];
                size *= d;
                s->dims[i] = d;
        }

        s->rank      = rank;
        s->ref_count = 1;
        s->size      = size;

        return s;
}

struct shape_t* spIncRef(struct shape_t* p)
{
        p->ref_count++;
        return p;
}

struct shape_t* spDecRef(struct shape_t* p)
{
        if (--(p->ref_count) == 0) {
                // free the shape.
                free(p);
                return NULL;
        }
        return p;
}
