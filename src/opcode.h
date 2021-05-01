enum opcode_t {
        // Option.
        //   - Option can be optional. If so, the NULL is provided and it
        //     matches the default setting.
        //   - If Option.mode is used only, then macros should be provided, so
        //     the caller can have some readability.
        //   - If Option union {i,f,etc} is used, to avoid the sitution the
        //     caller forgets to update them, option.mode be must set in an
        //     explicit way so the check can be performed.

        // --------------------------------------------------------------------
        // Element ops
        // --------------------------------------------------------------------
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
        //     - For this case, as rhs is NULL, special value in mode will not
        //       be needed.
        //
        // In-Place:
        //   - all operand handles can be used as destination.
        OP_ADD,
        OP_MUL,
        OP_MINUS,
        OP_MAX,
        OP_CMPL,  // Compare large.

        // --------------------------------------------------------------------
        // Matmul
        // --------------------------------------------------------------------
        //
        // Data Types:
        //   - only F32.
        //
        // Shapes:
        //   - only rank 2 operands.
        //
        // Option:
        //   - opt could be NULL, or opt.mode == 0, i.e., OPT_MATMUL_TRANS_NOT
        //   - mode == 2 means trans_lhs, i.e., OPT_MATMUL_TRANS_LHS
        //   - mode == 1 means trans_rhs, i.e., OPT_MATMUL_TRANS_RHS
        //   - other values of modes are invalid.
        //
        // In-Place:
        //   - dst must be unique.
        OP_MATMUL,

        // --------------------------------------------------------------------
        // Reduction
        // --------------------------------------------------------------------
        //
        // Data Types:
        //   - only F32.
        //
        // Option (required):
        //   - opt.mode value table
        //
        //     | v | reduction op | macro                 |
        //     | 0 | sum          | OPT_SET_REDUCTION_SUM |
        //
        //   - opt.mode I bit (set after opt.mode)
        //     - if set, then opt.i specifies the axis. Use
        //       OPT_SET_REDUCTION_AXIS.
        //     - otherwise, opt.i == 0
        OP_REDUCE,

        OP_RNG,  // used .rng_seed for seed, mode for distribution.

        // --------------------------------------------------------------------
        // Softmax crossentropy with logits loss
        // --------------------------------------------------------------------
        //
        // Data Types:
        //   - only F32.
        //
        // Option:
        //   - opt could be NULL.
        //   - opt.mode I bit
        //     - if set then opt.i for tensor handle of grad w.r.t. o_i.
        //       Use OPT_SET_GRAD_TENSOR_HANDLER.
        //     - otherwise, opt.i == 0
        OP_LS_SCEL
};

// --- opt bits.
#define OPT_MODE_I_BIT          0x10000
#define OPT_MODE_GET_I_BIT(opt) (((opt).mode) & OPT_MODE_I_BIT)

// --- common macros
// --- matmul
#define OPT_MATMUL_TRANS_NOT 0
#define OPT_MATMUL_TRANS_LHS 2
#define OPT_MATMUL_TRANS_RHS 1

// --- reduction
#define OPT_SET_REDUCTION_SUM(opt) ((opt).mode = 0, (opt).i = 0)

#define OPT_SET_REDUCTION_AXIS(opt, axis) \
        ((opt).mode |= OPT_MODE_I_BIT, (opt).i = (axis))

// --- loss
#define OPT_SET_GRAD_TENSOR_HANDLER(opt, td) \
        ((opt).mode |= OPT_MODE_I_BIT, (opt).i = (td))
