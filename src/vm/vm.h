#ifndef VM_H_
#define VM_H_

#include <stdlib.h>  // sizt_t

#include "base/error.h"

struct vm_t {
        size_t size_used;
};

extern error_t vmLaunch(struct vm_t*);
extern float   vmComsumedSizeInMB(struct vm_t*);

#endif
