#ifndef MLVM_COMPUTATION_INSTRUCTION_
#define MLVM_COMPUTATION_INSTRUCTION_

#include "mlvm/Computation/OpCode.h"
#include "mlvm/Computation/TensorLike.h"

#include <memory>
#include <string>
#include <vector>

namespace mlvm::computation {

class Function;

class Instruction {
 public:
  Instruction(std::string name, OpCode op, std::vector<TensorLike*>&& inputs,
              Function* parentFunc);

 public:
  std::string name() const { return name_; }

  OpCode opCode() const { return opCode_; }

  TensorLike* const getOutput(int i) { return outputs_[i].get(); }
  int outputsCount() const { return outputs_.size(); }

  std::string string() const;

 private:
  std::string name_;
  OpCode opCode_;
  std::vector<TensorLike*> inputs_;
  std::vector<std::unique_ptr<TensorLike>> outputs_;

  Function* parentFunc_;
};

}  // namespace mlvm::computation

#endif
