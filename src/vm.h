#ifndef VM_H_
#define VM_H_

#include <stdint.h>

#include "adt/sds.h"
#include "base/error.h"
#include "rng/srng64.h"

// -----------------------------------------------------------------------------
// data structures.
// -----------------------------------------------------------------------------

typedef float float32_t;

enum data_t {
        F32,  // float32_t
        I32,  // int32_t
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

#define VM_UNUSED -1

struct vm_t;  // forward def.

enum opcode_t {
        // Element ops
        //
        // Data Types:
        //   - only F32.
        //
        // Shapes:
        //   - shapes are equal match.
        //   - lhs shape is multiple of rhs shape (broadcasting).
        //   - rhs is scalar, i.e., [1].
        //   - rhs is NULL; uses .f for rhs scalar [F32].
        //
        // Option:
        //   - opt must be NULL if .f is not used.
        //
        // In-Place:
        //   - all operand handles can be used as destination.
        OP_ADD,
        OP_MUL,
        OP_MINUS,

        // Matmul
        //
        // Data Types:
        //   - only F32.
        //
        // Shapes:
        //   - only rank 2 operands.
        //
        // Option:
        //   - opt must be NULL.
        //
        // In-Place:
        //   - dst must be unique.
        OP_MATMUL,

        OP_REDUCE,
        OP_RNG,  // used .rng_seed for seed, mode for distribution.
};

struct opopt_t {
        int mode;  // distribution mode for rng.
        union {
                const struct srng64_t* rng_seed;  // unowned.
                float32_t              f;
                int32_t                i;
        };
};

// -----------------------------------------------------------------------------
// apis for vm.
// -----------------------------------------------------------------------------

struct vm_t* vmNew();
void         vmFree(struct vm_t*);
error_t      vmExec(struct vm_t* vm, enum opcode_t, const struct opopt_t* opt,
                    int dst, int lhs, int rhs);

// -----------------------------------------------------------------------------
// apis for tensors. / tensor.c
// -----------------------------------------------------------------------------

int     vmTensorNew(struct vm_t*, enum data_t, struct shape_t*);
error_t vmTensorFree(struct vm_t*, int t);

error_t vmTensorInfo(struct vm_t*, int t, _mut_ enum data_t*,
                     _mut_ struct shape_t**);
error_t vmTensorData(struct vm_t*, int t, _mut_ void** data);
error_t vmTensorSwap(struct vm_t*, int t, _mut_ void** data);
void    vmTensorDump(sds_t* s, struct vm_t*, int t);

// -----------------------------------------------------------------------------
// apis for shapes. / shape.c
// -----------------------------------------------------------------------------

struct shape_t* vmShapeNew(struct vm_t* vm, int rank, int* dims);
struct shape_t* spNew(int rank, int* dims);
struct shape_t* spIncRef(struct shape_t*);
struct shape_t* spDecRef(struct shape_t*);

#endif
