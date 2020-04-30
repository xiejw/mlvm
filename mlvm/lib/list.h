#ifndef MLVM_LIB_LIST_H_
#define MLVM_LIB_LIST_H_

#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "mlvm/lib/types.h"

typedef struct {
  mlvm_size_t size;
  mlvm_size_t cap;
} list_base_t;

#define list_t(T)     \
  struct {            \
    list_base_t base; \
    T*          data; \
  }

#define list_init(l)             \
  (memset((l), 0, sizeof(*(l))), \
   list_allocate_(&(l)->base, (void**)&(l)->data, sizeof(*(l)->data)))

#define list_deinit(l) (free((l)->data))

#define list_deinit_with_deleter(l, deleter)  \
  do {                                        \
    mlvm_size_t index, size = (l)->base.size; \
    for (index = 0; index < size; index++) {  \
      deleter((l)->data[index]);              \
    }                                         \
    free((l)->data);                          \
  } while (0)

#define list_size(l) ((l)->base.size)

#define list_set(l, i, v) (*((l)->data + i) = v)

#define list_append(l, v)                                                    \
  ((l)->base.cap != (l)->base.size + 1                                       \
       ? (void)0                                                             \
       : list_may_grow_(&(l)->base, (void**)&(l)->data, sizeof(*(l)->data)), \
   *((l)->data + (l)->base.size++) = (v))

extern void list_allocate_(list_base_t* base, void** data, int vsize);
extern void list_may_grow_(list_base_t* base, void** data, int vsize);

/* Common data structures. */
typedef list_t(int) list_int_t;

#endif
