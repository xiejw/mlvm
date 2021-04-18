#include "vm.h"

#include <stdlib.h>

// consider to use tree so the lookup is faster.
struct shape_list_t {
        struct shape_t*      s;
        struct shape_list_t* next;
};

static struct shape_list_t* root = NULL;

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

        // add to root
        struct shape_list_t* new_node = malloc(sizeof(struct shape_list_t));
        new_node->s                   = s;
        if (root == NULL) {
                new_node->next = NULL;
                root           = new_node;
        } else {
                new_node->next = root;
                root           = new_node;
        }

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
                // delete p from root.
                assert(root != NULL);
                struct shape_list_t* cur = root;
                struct shape_list_t* pre = NULL;
                assert(root != NULL);

                while (1) {
                        if (cur->s == p) {
                                if (pre == NULL) {  // cur is root.
                                        root = cur->next;
                                        free(cur);
                                        break;
                                } else {
                                        pre->next = cur->next;
                                        free(cur);
                                        break;
                                }
                        } else {
                                pre = cur;
                                cur = cur->next;
                                assert(cur != NULL);
                        }
                }

                // free the shape.
                free(p);
                return NULL;
        }
        return p;
}

void spFreeAll() {}
