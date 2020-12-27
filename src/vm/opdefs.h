#ifndef OPDEFS_H_
#define OPDEFS_H_

// clang-format off
#define _NO_OPERAND 0, { 0 }
#define _ONE_UINT16 1, { 2 }
// clang-format on

// clang-format off
static opdef_t opDefs[OP_END] = {

    // Loads constant object from Program into stack.
    //
    // Operand: (uint16) object index.
    // Stack  : push the object to the top.
    {"OP_CONST", _ONE_UINT16},

    // Pops out top item on stack.
    //
    // Operand: no.
    // Stack  : pop the object from the top.
    {"OP_POP", _NO_OPERAND},

    // OP_LOAD,
    // OP_MOVE,
    // OP_STORE,

    // OP_RNG,
    // OP_RNGT,
    // OP_RNGS,
};
// clang-format on

#endif
