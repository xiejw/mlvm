// folder internal only.
#ifndef OPDEFS_H_
#define OPDEFS_H_

// clang-format off
#define _NO_OPERAND 0, { 0 }
#define _ONE_UINT16 1, { 2 }

static struct opdef_t opDefs[] = {
    {"OP_HALT", _NO_OPERAND}, // halts the execution.
};
// clang-format on

static int opCount = sizeof(opDefs) / sizeof(struct opdef_t);

#endif
