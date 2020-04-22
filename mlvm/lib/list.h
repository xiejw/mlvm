#ifndef MLVM_CONTAINER_LIST_H_
#define MLVM_CONTAINER_LIST_H_

#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
  uint64_t size;
  uint64_t cap;
} list_base_t;

#define list_t(T)     \
  struct {            \
    list_base_t base; \
    T*          ref;  \
  }

#define list_init(l) (memset((l), 0, sizeof(*(l))))

#define list_deinit(l) (free((l)->ref))

#define list_size(l) ((l)->base.size)

#define list_get(l, i) (*((l)->ref + i))

#define list_set(l, i, v) (*((l)->ref + i) = v)

#define list_append(l, v)                                            \
  (list_may_grow_(&(l)->base, (void**)&(l)->ref, sizeof(*(l)->ref)), \
   *((l)->ref + (l)->base.size++) = (v))

extern void list_may_grow_(list_base_t* base, void** ref, int vsize);

/* Common data structures. */
typedef list_t(int) list_int_t;

#endif
