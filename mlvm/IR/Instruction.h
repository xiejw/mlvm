#ifndef MLVM_IR_INSTRUCTION_H_
#define MLVM_IR_INSTRUCTION_H_

#include <vector>

#include "mlvm/IR/Tensor.h"

namespace mlvm {

enum class OpType { Add };

class Instruction {
 public:
  Instruction(std::string name, OpType op, std::vector<Tensor*> inputs)
      : name_{std::move(name)}, op_{op}, inputs_{std::move(inputs)} {};

  // TODO: Use outputs.
  void BuildOutputs();

  std::string_view name() const { return name_; }

  Tensor* outputAt(int i) const { return outputs_[i].get(); }

  std::string debugString() const;

 private:
  std::string name_;
  OpType op_;
  std::vector<Tensor*> inputs_;  // unowned.
  std::vector<std::unique_ptr<Tensor>> outputs_ = {};
};

}  // namespace mlvm

#endif
