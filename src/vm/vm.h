#ifndef VM_H_
#define VM_H_

#include <stdlib.h>  // sizt_t

#include "adt/vec.h"
#include "base/error.h"
#include "object.h"
#include "opcode.h"

struct vm_t {
        size_t size_used;

        // internal fields.
        struct obj_t* base;
        struct obj_t* top;
        struct obj_t* stack;
};

extern struct vm_t* vmNew(void);
extern void         vmFree(struct vm_t* vm);

extern error_t vmLaunch(struct vm_t* vm, vec_t(code_t));
extern float   vmComsumedSizeInMB(struct vm_t* vm);

#endif
