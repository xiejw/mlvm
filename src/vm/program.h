#ifndef PROGRAM_H_
#define PROGRAM_H_

#include "object.h"
#include "opcode.h"
#include "vec.h"

typddef

    typedef struct {
        vec_t(obj_t*) data;
        // vect(codet) instructions;
} program_tt;

// programt* pgCreate();
// void      pgFree(programt* pg);

#endif
