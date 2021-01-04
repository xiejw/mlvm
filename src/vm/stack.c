#include "stack.h"

#include <stdio.h>
#include <string.h>

#include "adt/vec.h"
#include "object.h"
#include "opcode.h"

#define STACK_INIT_SIZE 256
static obj_t *base  = NULL;
static obj_t *top   = NULL;
static obj_t *stack = NULL;

static opcode_t *pc = NULL;

void stackInit() {
  if (stack != NULL) free(stack);
  stack = malloc(STACK_INIT_SIZE * sizeof(obj_t));
  base  = stack;
  top   = stack;
}

error_t vmExec(vec_t(opcode_t) code) {
  opcode_t op;
  pc = code;

  switch (op = *pc++) {
    case OP_HALT:
      printf("halt\n");
      return OK;
    default:
      return errFatalAndExit("unsupported opcode: %d", op);
  }
}
