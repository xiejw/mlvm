#include "mlvm/Computation/Instruction.h"

#include <sstream>

namespace mlvm::computation {

Instruction::Instruction(std::string name, OpCode op,
                         std::vector<TensorLike*>&& inputs)
    : name_{std::move(name)}, opCode_{op}, inputs_{inputs}, outputs_{} {}

std::string Instruction::string() const {
  std::stringstream str;

  // Name;
  str << name_ << " (";

  // Inputs;
  int inputs_size = inputs_.size();
  for (int i = 0; i < inputs_size; ++i) {
    auto& t = inputs_[i];
    str << "`" << t->name() << "`" << t->shape().string();

    if (i != inputs_size - 1) str << ", ";
  }
  str << ") -> ()";
  return str.str();
}

}  // namespace mlvm::computation
