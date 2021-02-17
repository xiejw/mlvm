#ifndef OPCODE_H_
#define OPCODE_H_

// #include "adt/sds.h"
#include "adt/vec.h"
#include "base/defs.h"

#define OPCODE_MAX_NUM_OPERANDS 1

// -----------------------------------------------------------------------------
// opcode.
// -----------------------------------------------------------------------------

enum opcode_t {
        OP_FILL,
};

struct opdef_t {
        char* name;
        int   allow_grad;  // if false, grad cannot flow.
        int   num_operands;
        int   num_outputs;
};

typedef int vm_handle_t;

struct op_record_t {
        enum opcode_t code;
        vec_t(vm_handle_t) operands;
        vec_t(vm_handle_t) outputs;
        void* option;
};

struct op_tape_t {
        vec_t(op_record_t*) records;
};

// -----------------------------------------------------------------------------
// prototypes.
// -----------------------------------------------------------------------------

// extern error_t opLookup(enum opcode_t c, _mut_ struct opdef_t** def);
// extern error_t opMake(_mut_ vec_t(code_t) * code, enum opcode_t c, ...);
// extern error_t opDump(_mut_ sds_t* buf, code_t* code, int size, char*
// prefix);

#endif
