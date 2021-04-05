#ifndef OP_H_
#define OP_H_

#include "vm.h"

error_t vmOpAddF32(struct tensor_t* dst, struct tensor_t* t1,
                   struct tensor_t* t2);

error_t vmOpMulF32(struct tensor_t* dst, struct tensor_t* t1,
                   struct tensor_t* t2);

error_t vmOpMinusF32(struct tensor_t* dst, struct tensor_t* t1,
                     struct tensor_t* t2);

error_t vmOpcRngF32(struct tensor_t* dst, int mode,
                    const struct srng64_t* seed);

#endif
