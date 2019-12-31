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
  TensorLike* const getOutputs(int i) { return outputs_[i].get(); };

 public:
  OpCode opCode() const { return opCode_; }

  std::string string() const;

 private:
  friend class Function;

  Instruction(std::string name, OpCode op, std::vector<TensorLike*>&& inputs,
              Function* parentFunc);

 private:
  std::string name_;
  OpCode opCode_;

  std::vector<TensorLike*> inputs_;
  std::vector<std::unique_ptr<TensorLike>> outputs_;

  Function* parentFunc_;
};

}  // namespace mlvm::computation

#endif
