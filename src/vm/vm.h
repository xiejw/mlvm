#ifndef VM_H_
#define VM_H_

#include <stdlib.h>  // sizt_t

#include "adt/vec.h"
#include "base/error.h"
#include "object.h"
#include "opcode.h"

typedef int vm_handle_t;  // zero means OOM.

struct vm_t {
        size_t size_used;
        void*  handles;  // opaque.
};

extern struct vm_t* vmNew(void);
extern void         vmFree(struct vm_t* vm);

extern vm_handle_t vmAllocTensor(struct vm_t* vm, int rank, int dims[]);
extern error_t     vmDeallocTensor(struct vm_t* vm, vm_handle_t);

extern void vmReset(struct vm_t* vm);

// extern error_t     vmRead(struct vm_t* vm, vm_handle_t, obj_float_t* dst);
// extern error_t     vmWrite(struct vm_t* vm, vm_handle_t, obj_float_t* src);
// extern error_t     vmLaunch(struct vm_t* vm, vec_t(code_t),
//                             vec_t(struct obj_tensor_t*) * outputs);
extern float vmComsumedSizeInMiB(struct vm_t* vm);

#endif
