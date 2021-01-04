#include "stack.h"

#include <string.h>

#include "object.h"

#define STACK_INIT_SIZE 256
static obj_t *pc   = NULL;
static obj_t *base = NULL;
static obj_t *top  = NULL;

void stackInit() {
  if (pc != NULL) free(pc);
  pc   = malloc(STACK_INIT_SIZE * sizeof(obj_t));
  base = pc;
  top  = pc;
}
