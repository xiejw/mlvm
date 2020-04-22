#include "mlvm/lib/list.h"

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
