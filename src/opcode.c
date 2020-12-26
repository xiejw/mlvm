#include "opcode.h"

#include <stdarg.h>
#include <stdio.h>  //del

// clang-format off
#define _NO_OPERAND 0, { 0 }
#define _ONE_UINT16 1, { 2 }
// clang-format on

static opdeft opDefs[OP_END] = {
    // clang-format off
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
    // clang-format on

    // OP_LOAD,
    // OP_MOVE,
    // OP_STORE,

    // OP_RNG,
    // OP_RNGT,
    // OP_RNGS,
};

#define _BIGENDIAN_PUT_UINT16(code, x)     \
  do {                                     \
    vecPushBack((code), (char)((x) >> 8)); \
    vecPushBack((code), (char)(x));        \
  } while (0)

errort opLookup(opcodet c, opdeft** def) {
  if (c >= 0 && c < OP_END) {
    *def = &opDefs[c];
    return OK;
  }
  return ENOT_FOUND;
}

errort opMake(opcodet c, vect(codet) * code, ...) {
  if (c >= 0 && c < OP_END) {
    opdeft* def      = &opDefs[c];
    int     num_args = def->num_operands;
    vecPushBack(*code, (codet)c);

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
            printf("unsupported width.\n");  // error message.
            return EUNSPECIFIED;
        }
      }

      va_end(ap);
    }
    return OK;
  }

  return ENOT_FOUND;
}
