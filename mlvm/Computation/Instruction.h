#ifndef MLVM_COMPUTATION_INSTRUCTION_
#define MLVM_COMPUTATION_INSTRUCTION_

#include "mlvm/Computation/OpCode.h"
#include "mlvm/Computation/TensorLike.h"

#include <vector>

namespace mlvm::computation {

class Function;

class Instruction {
 public:
  const TensorLike& getOutputs(int i) { return outputs_[i]; };

 public:
  OpCode opCode() const { return opCode_; }

 private:
  friend class Instruction;

  Instruction(OpCode op, std::vector<TensorLike*>&& inputs);

 private:
  OpCode opCode_;
  std::vector<TensorLike*> inputs_;
  std::vector<TensorLike> outputs_;
};

}  // namespace mlvm::computation

#endif
