#ifndef OPCODE_H_
#define OPCODE_H_

#include "mlvm.h"
#include "vec.h"

#define OPCODE_MAX_NUM_OPERANDS 1

typedef enum {
  OP_CONST,
  OP_POP,

  // OP_LOAD,
  // OP_MOVE,
  // OP_STORE,

  // OP_RNG,
  // OP_RNGT,
  // OP_RNGS,

  OP_END  // unused
} opcodet;

typedef struct {
  char* name;
  int   num_operands;
  int   widths[OPCODE_MAX_NUM_OPERANDS];
} opdeft;

typedef char codet;

// -----------------------------------------------------------------------------
// Prototypes.
// -----------------------------------------------------------------------------

errort opLookup(opcodet c, _mut_ opdeft** def);
errort opMake(opcodet c, _mut_ vect(codet) * code, ...);

#endif
