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
        // Option:
        //   - opt must be NULL.
        //
        // In-Place:
        //   - dst must be unique.
        OP_MATMUL,

        OP_REDUCE,
        OP_RNG,     // used .rng_seed for seed, mode for distribution.
        OP_LS_SCEL  // softmax crossentropy with logits loss
};
