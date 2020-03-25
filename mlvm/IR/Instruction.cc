#include "mlvm/IR/Instruction.h"

#include <sstream>

#include <absl/strings/str_cat.h>

namespace mlvm {

std::string Instruction::debugString() const {
  std::stringstream ss{};
  switch (op_) {
    case OpType::Add:
      ss << "Add";
      break;
    default:
      ss << "Unknown Op";
  }
  ss << " `" << name_ << "`";
  ss << " (" << inputs_.size() << " inputs and " << outputs_.size()
     << " outputs)";
  return ss.str();
}

void Instruction::BuildOutputs() {
  auto result = new OutputTensor{absl::StrCat("%o_{", name_, "}"), this, 0};
  outputs_.emplace_back(result);
};

}  // namespace mlvm
