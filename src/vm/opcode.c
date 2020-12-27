#include "vm/opcode.h"

#include <stdarg.h>

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

#define _BIGENDIAN_PUT_UINT16(code, x)     \
  do {                                     \
    vecPushBack((code), (char)((x) >> 8)); \
    vecPushBack((code), (char)(x));        \
  } while (0)

error_t opLookup(opcode_t c, opdef_t** def) {
  if (c >= 0 && c < OP_END) {
    *def = &opDefs[c];
    return OK;
  }
  return errNewWithNote(ENOTEXIST, "opcode does not exist: %d", c);
}

error_t opMake(opcode_t c, vec_t(code_t) * code, ...) {
  if (c >= 0 && c < OP_END) {
    opdef_t* def      = &opDefs[c];
    int      num_args = def->num_operands;
    vecPushBack(*code, (code_t)c);

    // Handles the operands.
    if (num_args > 0) {
      va_list ap;
      va_start(ap, code);

      for (int i = 0; i < num_args; i++) {
        int operand = va_arg(ap, int);
        switch (def->widths[i]) {
          case 2:
            _BIGENDIAN_PUT_UINT16(*code, operand);
            break;
          default:
            return errNewWithNote(ENOTIMPL, "unsupported width for code: %d",
                                  def->widths[i]);
        }
      }

      va_end(ap);
    }
    return OK;
  }

  return errNewWithNote(ENOTEXIST, "opcode does not exist: %d", c);
}
