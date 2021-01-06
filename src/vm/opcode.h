#ifndef OPCODE_H_
#define OPCODE_H_

#include "adt/vec.h"
#include "base/defs.h"

#define OPCODE_MAX_NUM_OPERANDS 1

// -----------------------------------------------------------------------------
// opcode.
// -----------------------------------------------------------------------------

enum opcode_t {
        OP_HALT,
};

struct opdef_t {
        char* name;
        int   num_operands;
        int   widths[OPCODE_MAX_NUM_OPERANDS];
};

typedef char code_t;

// -----------------------------------------------------------------------------
// prototypes.
// -----------------------------------------------------------------------------

extern error_t opLookup(enum opcode_t c, _mut_ struct opdef_t** def);
extern error_t opMake(enum opcode_t c, _mut_ vec_t(code_t) * code, ...);

#endif
