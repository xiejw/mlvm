#ifndef OPCODE_H_
#define OPCODE_H_

#include "adt/vec.h"
#include "base/defs.h"

#define OPCODE_MAX_NUM_OPERANDS 1

typedef enum {
  OP_HALT,
  OP_CONST,
  OP_POP,

  // OP_LOAD,
  // OP_MOVE,
  // OP_STORE,

  // OP_RNG,
  // OP_RNGT,
  // OP_RNGS,

  OP_END  // unused
} opcode_t;

typedef struct {
  char* name;
  int   num_operands;
  int   widths[OPCODE_MAX_NUM_OPERANDS];
} opdef_t;

typedef char code_t;

// -----------------------------------------------------------------------------
// prototypes.
// -----------------------------------------------------------------------------

extern error_t opLookup(opcode_t c, _mut_ opdef_t** def);
extern error_t opMake(opcode_t c, _mut_ vec_t(code_t) * code, ...);

#endif
