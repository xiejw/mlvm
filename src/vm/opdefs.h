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
    {"OP_HALT",         _NO_OPERAND},
    {"OP_PUSHBYTE",     _ONE_UINT8},
    {"OP_LOADGLOBAL",   _NO_OPERAND},
    {"OP_RETURN",       _ONE_UINT8},
};
// clang-format on

static int opTotalCount     = sizeof(opDefs) / sizeof(struct opdef_t);
static int opMaxCountOfChar = 15;  // change opDump if change here.

#endif
