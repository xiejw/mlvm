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

typedef list_t(int) list_int_t;

#define DEFAULT_LIST_SIZE 16

void list_may_grow_(list_base_t* base, void** ref, int vsize) {
  if (base->cap == 0) {
    printf("init\n");
    *ref      = malloc(DEFAULT_LIST_SIZE * vsize);
    base->cap = DEFAULT_LIST_SIZE;
  }

  if (base->cap == base->size + 1) {
    printf("grow\n");
    uint64_t new_cap = base->cap * 2;
    void*    new_ptr = malloc(new_cap * vsize);
    void*    old_ptr = ref;
    memcpy(new_ptr, old_ptr, base->cap);
    *ref = new_ptr;
    free(old_ptr);
  }
}

#endif
