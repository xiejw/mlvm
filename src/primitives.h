#ifndef OP_H_
#define OP_H_

#include "vm.h"

// -----------------------------------------------------------------------------
// element wise.
// -----------------------------------------------------------------------------
error_t vmOpAddF32(struct tensor_t* td, struct tensor_t*, struct tensor_t*);
error_t vmOpMulF32(struct tensor_t* td, struct tensor_t*, struct tensor_t*);
error_t vmOpMinusF32(struct tensor_t* td, struct tensor_t*, struct tensor_t*);
error_t vmOpMaxF32(struct tensor_t* td, struct tensor_t*, struct tensor_t*);
error_t vmOpCmpLF32(struct tensor_t* td, struct tensor_t*, struct tensor_t*);

error_t vmOpAddSF32(struct tensor_t* td, struct tensor_t* t1, float32_t s);
error_t vmOpMulSF32(struct tensor_t* td, struct tensor_t* t1, float32_t s);
error_t vmOpMinusSF32(struct tensor_t* td, struct tensor_t* t1, float32_t);
error_t vmOpMaxSF32(struct tensor_t* td, struct tensor_t* t1, float32_t);
error_t vmOpCmpLSF32(struct tensor_t* td, struct tensor_t* t1, float32_t);

// -----------------------------------------------------------------------------
// reduction.
// -----------------------------------------------------------------------------
error_t vmOpReduceF32(struct tensor_t* td, struct tensor_t* t1, int mode,
                      int axis);

// -----------------------------------------------------------------------------
// rng.
// -----------------------------------------------------------------------------
error_t vmOpRngF32(struct tensor_t* td, int mode, const struct srng64_t* seed);

// -----------------------------------------------------------------------------
// matmul.
// -----------------------------------------------------------------------------
error_t vmOpMatmulF32(struct tensor_t* td, struct tensor_t*, struct tensor_t*,
                      int trans_lhs, int trans_rhs);

// -----------------------------------------------------------------------------
// loss.
// -----------------------------------------------------------------------------
error_t vmOpLossSCELF32(struct tensor_t* td, struct tensor_t* y,
                        struct tensor_t* o, struct tensor_t* optional_g);

#endif
