#ifndef MLVM_COMPUTATION_INSTRUCTION_
#define MLVM_COMPUTATION_INSTRUCTION_

#include "mlvm/Computation/OpCode.h"
#include "mlvm/Computation/TensorLike.h"

#include <string>
#include <vector>

namespace mlvm::computation {

class Instruction {
 public:
  const TensorLike& getOutputs(int i) { return outputs_[i]; };

 public:
  OpCode opCode() const { return opCode_; }

  std::string string() const {return std::string{"ins"} + name_; }

 private:
  friend class Function;

  Instruction(std::string name, OpCode op, std::vector<TensorLike*>&& inputs);

 private:
  std::string name_;
  OpCode opCode_;
  std::vector<TensorLike*> inputs_;
  std::vector<TensorLike> outputs_;
};

}  // namespace mlvm::computation

#endif
