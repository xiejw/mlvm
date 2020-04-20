#ifndef MLVM_IR_TENSOR_H_
#define MLVM_IR_TENSOR_H_

#define < stdint.h>

typedef struct {
  uint64_t  rank;  /* ==0 means scalar. */
  uint64_t* shape; /* size is `rank` above. */
}

#endif
