#include "mlvm/IR/Instruction.h"

#include <sstream>

#include <absl/strings/str_cat.h>

namespace mlvm::IR {

std::string Instruction::debugString() const {
  std::stringstream ss{};
  // Push type.
  switch (op_) {
    case OpType::Add:
      ss << "Add";
      break;
    default:
      ss << "Unknown Op";
  }
  // Push name.
  ss << " `" << name_ << "`";
  // Push operands.
  ss << " (";
  for (decltype(inputs_.size()) i = 0; i < inputs_.size(); ++i) {
    if (i != 0) ss << ", ";
    ss << inputs_[i]->debugString();
  }
  ss << ") -> (";
  // Push outputs
  for (decltype(outputs_.size()) i = 0; i < outputs_.size(); ++i) {
    if (i != 0) ss << ", ";
    ss << outputs_[i]->debugString();
  }
  ss << ")";
  return ss.str();
}

Status Instruction::buildOutputs() {
  const int index = 0;
  auto result = new OutputTensor{absl::StrCat("%o_{", name_, ", ", index, "}"),
                                 this, index};
  outputs_.emplace_back(result);
  return Status::OK;
};

}  // namespace mlvm::IR