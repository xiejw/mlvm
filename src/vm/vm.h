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
        //        struct obj_t*         base;
        //        struct obj_t*         top;  // point to next value to use.
        //        struct obj_t*         stack;
        //        struct obj_tensor_t** handles;
};

extern struct vm_t* vmNew(void);
extern void         vmFree(struct vm_t* vm);

// extern vm_handle_t vmAllocTensor(struct vm_t* vm, int rank, int dims[]);
// extern error_t     vmDeallocTensor(struct vm_t* vm, vm_handle_t);
// extern error_t     vmRead(struct vm_t* vm, vm_handle_t, obj_float_t* dst);
// extern error_t     vmWrite(struct vm_t* vm, vm_handle_t, obj_float_t* src);
// extern error_t     vmLaunch(struct vm_t* vm, vec_t(code_t),
//                             vec_t(struct obj_tensor_t*) * outputs);
extern float vmComsumedSizeInMiB(struct vm_t* vm);

#endif
