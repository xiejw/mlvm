#include "opcode.h"

#include <stdarg.h>

#include "opdefs.h"

#define _BIGENDIAN_PUT_UINT16(code, x)                 \
        do {                                           \
                vecPushBack((code), (char)((x) >> 8)); \
                vecPushBack((code), (char)(x));        \
        } while (0)

error_t opLookup(enum opcode_t c, struct opdef_t** def)
{
        if (c >= 0 && c < opCount) {
                *def = &opDefs[c];
                return OK;
        }
        return errNewWithNote(ENOTEXIST, "opcode does not exist: %d", c);
}

error_t opMake(enum opcode_t c, vec_t(code_t) * code, ...)
{
        if (c >= 0 && c < opCount) {
                struct opdef_t* def      = &opDefs[c];
                int             num_args = def->num_operands;
                vecPushBack(*code, (code_t)c);

                // Handles the operands.
                if (num_args > 0) {
                        va_list ap;
                        va_start(ap, code);

                        for (int i = 0; i < num_args; i++) {
                                int operand = va_arg(ap, int);
                                switch (def->widths[i]) {
                                        case 2:
                                                _BIGENDIAN_PUT_UINT16(*code,
                                                                      operand);
                                                break;
                                        default:
                                                return errNewWithNote(
                                                    ENOTIMPL,
                                                    "unsupported width for "
                                                    "code: %d",
                                                    def->widths[i]);
                                }
                        }

                        va_end(ap);
                }
                return OK;
        }

        return errNewWithNote(ENOTEXIST, "opcode does not exist: %d", c);
}
