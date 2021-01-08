#ifndef OPDEFS_H_
#define OPDEFS_H_

// -----------------------------------------------------------------------------
// opcode definitions. internal only.
// -----------------------------------------------------------------------------

// clang-format off
#define _NO_OPERAND 0, { 0 }
#define _ONE_UINT16 1, { 2 }
#define _ONE_UINT8  1, { 1 }

static struct opdef_t opDefs[] = {
    {"OP_HALT",      _NO_OPERAND}, // halt the execution.
    {"OP_PUSHBYTE",  _ONE_UINT8},  // push a byte to stack top.
};
// clang-format on

static int opCount = sizeof(opDefs) / sizeof(struct opdef_t);

#endif
