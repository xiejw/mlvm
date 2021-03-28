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

struct tensor_t {
        enum data_t     dtype : 7;
        int             used : 1;
        struct shape_t* shape;
        void*           data;
};

#define MAX_TENSOR_COUNT 128

struct vm_t {
        // consider to use pages.
        struct tensor_t handles[MAX_TENSOR_COUNT];
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
// apis for tensors. / tensor.c
// -----------------------------------------------------------------------------

int     vmNewT(struct vm_t*, enum data_t, struct shape_t*);
error_t vmFreeT(struct vm_t*, int);

error_t vmFetchMetadata(struct vm_t*, int handle, enum data_t*,
                        struct shape_t**);
error_t vmFetchData(struct vm_t*, int handle, void** data);

// -----------------------------------------------------------------------------
// apis for shapes. / shape.c
// -----------------------------------------------------------------------------

struct shape_t* spNew(int rank, int* dims);
struct shape_t* spIncRef(struct shape_t*);
struct shape_t* spDecRef(struct shape_t*);

// -----------------------------------------------------------------------------
// internal apis.
// -----------------------------------------------------------------------------

void vmFreeHandle(struct tensor_t* t);  // tensor.c

#endif
