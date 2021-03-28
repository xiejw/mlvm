#ifndef VM_H_
#define VM_H_

#include <stdint.h>

#include "base/error.h"

// -----------------------------------------------------------------------------
// data structures.
// -----------------------------------------------------------------------------

enum data_t {
        F32,
        I32,
};

struct shape_t {
        int      rank;
        uint64_t size;
        int      ref_count;
        int      dims[];
};

struct vm_t {
        int handles[128];
};

enum opcode_t {
        OP_ADD,
};

// -----------------------------------------------------------------------------
// apis for vm.
// -----------------------------------------------------------------------------

struct vm_t* vmNew();
void         vmFree();

int  vmExec(enum opcode_t, void* opt, int lhs, int rhs);
void vmSync();

// -----------------------------------------------------------------------------
// apis for tensors.
// -----------------------------------------------------------------------------

int     vmNewT(struct vm_t*, enum data_t, struct shape_t*);
error_t vmFreeT(int);

error_t vmFetchMetadata(struct vm_t*, int handle, enum data_t*,
                        struct shape_t**);
error_t vmFetchData(struct vm_t*, int handle, void** ptr);

// -----------------------------------------------------------------------------
// apis for shapes.
// -----------------------------------------------------------------------------

struct shape_t* spNew(int rank, int* dims);
void            spIncRef(struct shape_t*);
void            spDecRef(struct shape_t*);

#endif
