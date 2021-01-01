#ifndef PROGRAM_H_
#define PROGRAM_H_

#include "opcode.h"
#include "vec.h"
#include "object.h"

typddef

typedef struct {
  vec_t(obj_t*) data;
  // vect(codet) instructions;
} program_tt;

//programt* pgCreate();
//void      pgFree(programt* pg);

#endif
