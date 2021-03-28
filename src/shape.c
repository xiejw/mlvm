#include "vm.h"

#include <stdlib.h>

struct shape_t* spNew(int rank, int* dims)
{
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

void spIncRef(struct shape_t* p) { p->ref_count++; }

void spDecRef(struct shape_t* p)
{
        if (--(p->ref_count) == 0) {
                free(p);
        }
}
