#ifndef OPCODE_H_
#define OPCODE_H_

#define OPCODE_MAX_NUM_OPERANDS 1
#define OPCODE_MAX_CODE_LEN     OPCODE_MAX_NUM_OPERANDS + 1

typedef enum {
  OP_CONST,
  OP_POP,

  // OP_LOAD,
  // OP_MOVE,
  // OP_STORE,

  // OP_RNG,
  // OP_RNGT,
  // OP_RNGS,

  OP_END  // unused
} opcodet;

#endif
