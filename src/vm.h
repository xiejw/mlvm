#ifndef VM_H_
#define VM_H_

#include <stdint.h>

// eva
#include "adt/sds.h"
#include "base/error.h"
#include "rng/srng64.h"

// -----------------------------------------------------------------------------
// Data structures.
// -----------------------------------------------------------------------------

typedef float   f32_t;
typedef int32_t i32_t;

enum data_t {
        F32,  // f32_t
        I32,  // i32_t
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
        struct shape_t *shape;
        void           *data;
};

struct vm_t;  // forward def.

struct opopt_t {
        int mode;
        union {
                i32_t          i;
                f32_t          f;
                struct rng64_t r;
        };
};

// Enum opcode_t;
#include "op.h"

// -----------------------------------------------------------------------------
// APIs for vm.
// -----------------------------------------------------------------------------

struct vm_t *vmNew();
void         vmFree(struct vm_t *);
error_t      vmExec(struct vm_t *vm, enum opcode_t, const struct opopt_t *opt,
                    int dst, int lhs, int rhs);

// -----------------------------------------------------------------------------
// APIs for Batch Execution.
// -----------------------------------------------------------------------------
struct oparg_t {
        enum opcode_t  op;
        int            dst;
        int            t1;
        int            t2;
        int            has_opt;
        struct opopt_t opt;
};

error_t vmBatch(struct vm_t *vm, size_t arg_size, const struct oparg_t *);

// -----------------------------------------------------------------------------
// APIs for tensors.  / tensor.c
// -----------------------------------------------------------------------------

int     vmTensorNew(struct vm_t *, enum data_t, struct shape_t *);
error_t vmTensorFree(struct vm_t *, int t);

error_t vmTensorInfo(struct vm_t *, int t, _mut_ enum data_t *,
                     _mut_ struct shape_t **);
error_t vmTensorData(struct vm_t *, int t, _mut_ void **data);
error_t vmTensorSwap(struct vm_t *, int t, _mut_ void **data);
void    vmTensorDump(sds_t *s, struct vm_t *, int t);

// -----------------------------------------------------------------------------
// APIs for shapes.  / shape.c
// -----------------------------------------------------------------------------

struct shape_t *spNew(int rank, int *dims);
struct shape_t *spIncRef(struct shape_t *);
struct shape_t *spDecRef(struct shape_t *);

struct shape_t *vmShapeNew(struct vm_t *vm, int rank, int *dims);

// Macors for shapes.
#define R1S(vm, s1)     vmShapeNew(vm, 1, (int[]){(s1)});
#define R2S(vm, s1, s2) vmShapeNew(vm, 2, (int[]){(s1), (s2)});

#endif
