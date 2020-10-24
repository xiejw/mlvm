#include "opcode.h"

typedef int error_t;

typedef struct {
  char* name;
  int   num_operand;
  int   widths[1];  // operands width. maximum as 1 now.
} opdeft;

static opdeft opDefs[OP_END] = {
    // clang-format off
    // Loads constant object from Program into stack.
    //
    // Operand: (uint16) object index.
    // Stack  : push the object to the top.
    {"OP_CONST", 1, {2},},

    // Pops out top item on stack.
    //
    // Operand: no.
    // Stack  : pop the object from the top.
    {"OP_POP", 0, {0}},
    // clang-format on

    // OP_LOAD,
    // OP_MOVE,
    // OP_STORE,

    // OP_RNG,
    // OP_RNGT,
    // OP_RNGS,
};

#define OK         0
#define ENOT_FOUND -2

error_t opLookup(opcodet c, opdeft** def) {
  if (c < 0 || c >= OP_END) return ENOT_FOUND;
  *def = &opDefs[c];
  return OK;
}
