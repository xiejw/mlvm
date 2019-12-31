#include "mlvm/Computation/Instruction.h"

#include <sstream>

namespace mlvm::computation {

Instruction::Instruction(std::string name, OpCode op,
                         std::vector<TensorLike*>&& inputs,
                         Function* parentFunc)
    : name_{std::move(name)},
      opCode_{op},
      inputs_{inputs},
      outputs_{},
      parentFunc_{parentFunc} {
  assert(op == OpCode::Add);
  // Assert shape equal or compatible.
  //
  // TODO: generate name
  auto o = new TensorLike{"%o", inputs_[0]->shape(), parentFunc_, this};

  outputs_.emplace_back(o);
}

std::string Instruction::string() const {
  std::stringstream str;

  // Name;
  str << name_ << " (";

  // Inputs;
  {
    int inputs_size = inputs_.size();
    for (int i = 0; i < inputs_size; ++i) {
      auto& t = inputs_[i];
      str << "`" << t->name() << "`" << t->shape().string();

      if (i != inputs_size - 1) str << ", ";
    }
  }
  str << ") -> (";

  // Outputs;
  {
    int outputs_size = outputs_.size();
    for (int i = 0; i < outputs_size; ++i) {
      auto& t = outputs_[i];
      str << "`" << t->name() << "`" << t->shape().string();

      if (i != outputs_size - 1) str << ", ";
    }
  }
  str << ")";

  return str.str();
}

}  // namespace mlvm::computation
