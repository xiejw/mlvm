#ifndef PROGRAM_H_
#define PROGRAM_H_

#include "opcode.h"
#include "vec.h"

typedef struct {
  vect(codet) instructions;
} programt;

programt* pgCreate();
void      pgFree(programt* pg);

#endif
