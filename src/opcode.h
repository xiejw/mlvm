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
        OP_MAX,
        OP_CMPL,  // Compare large.

        // Matmul
        //
        // Data Types:
        //   - only F32.
        //
        // Shapes:
        //   - only rank 2 operands.
        //
        // Option: (see macros below OPT_MATMUL_TRANS_*)
        //   - opt could be NULL, or opt.mode == 0. This means no transpose.
        //   - mode == 2 means trans_lhs
        //   - mode == 1 means trans_rhs
        //   - other values of modes are invalid.
        //
        // In-Place:
        //   - dst must be unique.
        OP_MATMUL,

        OP_REDUCE,
        OP_RNG,  // used .rng_seed for seed, mode for distribution.

        // Softmax crossentropy with logits loss
        //
        // Data Types:
        //   - only F32.
        //
        // Option:
        //   - (optional) opt.i for tensor handle of grad w.r.t. o_i
        OP_LS_SCEL
};

// --- common macros
#define OPT_MATMUL_TRANS_NOT 0
#define OPT_MATMUL_TRANS_LHS 2
#define OPT_MATMUL_TRANS_RHS 1
