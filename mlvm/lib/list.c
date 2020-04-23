#include "mlvm/lib/list.h"

#include <assert.h>

#define DEFAULT_LIST_SIZE 16

#ifdef DEBUG_LIST
#define DEBUG_LIST_PRINTF(x) printf(x)
#else
#define DEBUG_LIST_PRINTF(x)
#endif

void list_allocate_(list_base_t* base, void** ref, int vsize) {
  assert(base->cap == 0);
  DEBUG_LIST_PRINTF("Init List\n");
  *ref      = malloc(DEFAULT_LIST_SIZE * vsize);
  base->cap = DEFAULT_LIST_SIZE;
}

void list_may_grow_(list_base_t* base, void** ref, int vsize) {
  assert(base->cap == base->size + 1);
  DEBUG_LIST_PRINTF("Grow List buffer \n");
  uint64_t new_cap = base->cap * 2;
  void*    new_ptr = malloc(new_cap * vsize);
  void*    old_ptr = *ref;
  memcpy(new_ptr, old_ptr, base->cap * vsize);
  *ref = new_ptr;
  free(old_ptr);
}
