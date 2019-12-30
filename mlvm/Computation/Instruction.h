#ifndef MLVM_COMPUTATION_INSTRUCTION_
#define MLVM_COMPUTATION_INSTRUCTION_

#include "mlvm/Computation/TensorLike.h"
#include "mlvm/Computation/OpCode.h"

#include <vector>

namespace mlvm::computation {


class Instruction {
 public:
  Instruction(OpCode op) : opCode_{op} {};

 public:
  const TensorLike& getOutputs(int i) { return outputs_[i]; };

 private:
  OpCode opCode_;
  std::vector<TensorLike*> inputs_;
  std::vector<TensorLike> outputs_;
};


}  // namespace mlvm::computation

#endif
