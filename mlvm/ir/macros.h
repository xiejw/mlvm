#ifndef MLVM_IR_MACROS_H_
#define MLVM_IR_MACROS_H_

/* Utilities macros for internal usage. */

#include <assert.h>
#include <stdarg.h>

#define IR_MAX_NAME_SIZE 128

#define MLVM_IR_FILL_NAME(name, fmt)                         \
  {                                                          \
    va_list args;                                            \
    int     n;                                               \
    va_start(args, fmt);                                     \
    (name) = malloc(IR_MAX_NAME_SIZE * sizeof(char));        \
    n      = vsnprintf((name), IR_MAX_NAME_SIZE, fmt, args); \
    assert(n < IR_MAX_NAME_SIZE);                            \
    va_end(args);                                            \
  }

#endif
