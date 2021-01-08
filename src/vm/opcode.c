#include "opcode.h"

#include <stdarg.h>

#include "opdefs.h"

#define _BIGENDIAN_PUT_UINT8(code, x)             \
        do {                                      \
                vecPushBack((code), (code_t)(x)); \
        } while (0)

#define _BIGENDIAN_PUT_UINT16(code, x)                   \
        do {                                             \
                vecPushBack((code), (code_t)((x) >> 8)); \
                vecPushBack((code), (code_t)(x));        \
        } while (0)

// -----------------------------------------------------------------------------
// implementation.
// -----------------------------------------------------------------------------

error_t opLookup(enum opcode_t c, struct opdef_t** def)
{
        if (c < 0 || c >= opCount)
                return errNewWithNote(
                    ENOTEXIST, "opcode (%d) does not exist. total count %d", c,
                    opCount);

        *def = &opDefs[c];
        return OK;
}

error_t opMake(vec_t(code_t) * code, enum opcode_t c, ...)
{
        if (c < 0 || c >= opCount)
                return errNewWithNote(ENOTEXIST, "opcode does not exist: %d",
                                      c);

        struct opdef_t* def      = &opDefs[c];
        int             num_args = def->num_operands;
        vecPushBack(*code, (code_t)c);

        // Handles the operands.
        if (num_args == 0) goto done;

        va_list ap;
        va_start(ap, c);
        for (int i = 0; i < num_args; i++) {
                int operand = va_arg(ap, int);
                switch (def->widths[i]) {
                case 1:
                        _BIGENDIAN_PUT_UINT8(*code, operand);
                        break;
                case 2:
                        _BIGENDIAN_PUT_UINT16(*code, operand);
                        break;
                default:
                        return errNewWithNote(ENOTIMPL,
                                              "unsupported width for "
                                              "code: %d",
                                              def->widths[i]);
                }
        }
        va_end(ap);

done:
        return OK;
}
