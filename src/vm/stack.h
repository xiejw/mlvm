#ifndef STACK_H_
#define STACK_H_

#include "adt/vec.h"
#include "base/error.h"
#include "opcode.h"

error_t vmExec(vec_t(code_t));

#endif
