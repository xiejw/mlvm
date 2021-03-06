// -----------------------------------------------------------------------------
// Op code with spec.
//
// Not designed to be included in general. Use vm.h instead.
// -----------------------------------------------------------------------------
enum opcode_t {
        // --------------------------------------------------------------------
        // General Design Guidelines.
        // --------------------------------------------------------------------
        // Option.
        //   - Option can be optional. If so, the NULL is provided and it
        //     means the default setting.
        //
        //   - If Option.mode is used only, then macros is recommended, e.g.,
        //     OPT_MODE_I_BIT, etc. So the caller can have some readability.
        //
        //   - If Option union {i,f,r} is used, to avoid the sitution the
        //     caller forgets to update them, option.mode be must set in an
        //     explicit way so the check can be performed.

        // --------------------------------------------------------------------
        // Element ops
        // --------------------------------------------------------------------
        //
        // Operands:
        //   - both, if opt is NULL.
        //   - t1 only if opt is not NULL
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
        //   - opt could be NULL
        //   - if opt is not NULL, F_BIT must be set and .f specicies the
        //     scalar operand (the second operand).
        //
        // In-Place:
        //   - all operand handles can be used as destination.
        OP_ADD,
        OP_MUL,
        OP_MINUS,
        OP_DIVIDE,
        OP_MAX,
        OP_EQ,
        OP_CMPL,  // Compare Larger.

        // --------------------------------------------------------------------
        // Matmul
        // --------------------------------------------------------------------
        //
        // Operands:
        //   - both.
        //
        // Data Types:
        //   - only F32.
        //
        // Shapes:
        //   - only rank 2 operands.
        //
        // Option:
        //   - opt could be NULL. It means opt.mode == 0, i.e.,
        //     OPT_MATMUL_TRANS_NOT
        //
        //   - mode == 2 means trans_lhs, i.e., OPT_MATMUL_TRANS_LHS
        //   - mode == 1 means trans_rhs, i.e., OPT_MATMUL_TRANS_RHS
        //   - other values of modes are invalid.
        //
        // In-Place:
        //   - dst must be unique.
        OP_MATMUL,

        // --------------------------------------------------------------------
        // Arg
        // --------------------------------------------------------------------
        //
        // Operands:
        //   - t1 only.
        //
        // Data Types:
        //   - only F32.
        //
        // Shapes:
        //   - only rank 2.
        //
        // Option:
        //   - NULL
        //
        // In-Place:
        //   - No. dst is rank 1.
        OP_ARGMAX,

        // --------------------------------------------------------------------
        // Inverse Sqrt
        // --------------------------------------------------------------------
        //
        // Operands:
        //   - t1 only.
        //
        // Data Types:
        //   - only F32.
        //
        // Shapes:
        //   - any.
        //
        // Option (optional):
        //   - opt could be NULL. Defaults to no epsilon.
        //
        //   - If not NULL, then .f specicies the epsilon that
        //
        //       1 / sqrt(t1+epsilon)    if opt.mode = 0
        //       1 / (sqrt(t1)+epsilon)  if opt.mode = 1
        //
        // In-Place:
        //   - dst is allowed to be t1.
        OP_ISQRT,

        // --------------------------------------------------------------------
        // Reduction
        // --------------------------------------------------------------------
        //
        // Operands:
        //   - t1 only.
        //
        // Data Types:
        //   - only F32.
        //
        // Shapes:
        //   - any.
        //
        // Option (required):
        //   - opt.mode value table
        //
        //     | v | reduction op | macro                 |
        //     | 0 | sum          | OPT_SET_REDUCTION_SUM |
        //
        //   - opt.mode I bit
        //     - if set, then opt.i specifies the axis. Use
        //       OPT_SET_REDUCTION_SUM.
        //     - otherwise, opt.i must be 0
        //
        // In-Place:
        //   - not allowed.
        OP_REDUCE,

        // --------------------------------------------------------------------
        // Rng.
        // --------------------------------------------------------------------
        //
        // Operands:
        //   - no operands.
        //
        // Data Types:
        //   - only F32.
        //
        // Shapes:
        //   - any.
        //
        // Option (required):
        //   - opt.mode value table
        //
        //     | v | distribution | macro              |
        //     | 0 | std normal   | OPT_RNG_STD_NORMAL |
        //
        //   - opt.mode R bit
        //     must set .r to provide rng (seed).
        //
        // In-Place:
        //   - In-place anyway.
        OP_RNG,

        // --------------------------------------------------------------------
        // Fill.
        // --------------------------------------------------------------------
        //
        // Operands:
        //   - no operands.
        //
        // Data Types:
        //   - only F32.
        //
        // Shapes:
        //   - any.
        //
        // Option (optional):
        //   - NULL if fill with zero (optimized).
        //   - opt.f (F bit) to value to fill
        //
        // In-Place:
        //   - In-place anyway.
        OP_FILL,

        // --------------------------------------------------------------------
        // Softmax crossentropy with logits loss
        // --------------------------------------------------------------------
        //
        // Operands:
        //   - t1 is the label (must be valid distribution)
        //   - t2 is the logits (recommmended to be unnormalized).
        //
        // Data Types:
        //   - only F32.
        //
        // Shapes:
        //   - Rank 2 only for both.
        //
        // Option:
        //   - opt could be NULL.
        //   - if not NULL, opt.mode I bit must be set. Then opt.i for tensor
        //   handle of grad w.r.t. o_i.  Use OPT_SET_GRAD_TENSOR_HANDLER.
        //
        // In-Place:
        //   - Not allowed.
        OP_LS_SCEL
};

// -----------------------------------------------------------------------------
// Opt mask bits.
// -----------------------------------------------------------------------------
#define OPT_MODE_BIT_MASK       0xFF0000
#define OPT_MODE_UNMASK         0x00FFFF
#define OPT_MODE_I_BIT          0x10000
#define OPT_MODE_F_BIT          0x20000
#define OPT_MODE_R_BIT          0x40000
#define OPT_MODE_GET_I_BIT(opt) (((opt).mode) & OPT_MODE_I_BIT)
#define OPT_MODE_GET_F_BIT(opt) (((opt).mode) & OPT_MODE_F_BIT)
#define OPT_MODE_GET_R_BIT(opt) (((opt).mode) & OPT_MODE_R_BIT)

// -----------------------------------------------------------------------------
// Common macros
// -----------------------------------------------------------------------------
//
// --- Element wise ops
#define OPT_SET_SCALAR_OPERAND(opt, v) \
        ((opt).mode = OPT_MODE_F_BIT, (opt).f = (v))

// --- Matmul
#define OPT_MATMUL_TRANS_NOT 0
#define OPT_MATMUL_TRANS_LHS 2
#define OPT_MATMUL_TRANS_RHS 1

// --- Reduction
#define OPT_SET_REDUCTION_SUM(opt, axis) \
        ((opt).mode = 0 | OPT_MODE_I_BIT, (opt).i = (axis))

// --- Rng
#define OPT_RNG_STD_NORMAL 0

// --- Loss
#define OPT_SET_GRAD_TENSOR_HANDLER(opt, td) \
        ((opt).mode |= OPT_MODE_I_BIT, (opt).i = (td))
