#ifndef MLVM_IR_INSTRUCTION_H_
#define MLVM_IR_INSTRUCTION_H_

#include <memory>
#include <vector>

#include "mlvm/Foundation/Status.h"
#include "mlvm/IR/Tensor.h"

namespace mlvm::IR {

enum class OpType { Add };

class Instruction {
 public:
  Instruction(std::string name, OpType op, std::vector<Tensor*> inputs)
      : name_{std::move(name)}, op_{op}, inputs_{std::move(inputs)} {};

  std::string_view name() const { return name_; }

  std::string debugString() const;

  Status buildOutputs();

  Tensor* outputAt(int i) const { return outputs_[i].get(); }

 private:
  std::string name_;
  OpType op_;
  std::vector<Tensor*> inputs_;  // unowned.
  std::vector<std::unique_ptr<Tensor>> outputs_ = {};
};

}  // namespace mlvm::IR

#endif
