#ifndef OPCODE_H_
#define OPCODE_H_

#include "adt/sds.h"
#include "adt/vec.h"
#include "base/defs.h"

#define OPCODE_MAX_NUM_OPERANDS 1

// -----------------------------------------------------------------------------
// opcode.
// -----------------------------------------------------------------------------

enum opcode_t {
        OP_HALT,        // Halt machine.
        OP_PUSHBYTE,    // Push byte to stack. Code byte in next pc.
        OP_LOADGLOBAL,  // Get index from stack.
        OP_RETURN,      // Move values to base. Code count in next pc.
        OP_CFUNC,  // Call c func. Get func name from stack. Code return count
                   // in next pc.
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
extern error_t opMake(_mut_ vec_t(code_t) * code, enum opcode_t c, ...);
extern error_t opDump(_mut_ sds_t* buf, code_t* code, int size, char* prefix);

#endif
