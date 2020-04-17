#include "mlvm/IR/Function.h"

#include <sstream>

#include <absl/strings/str_cat.h>

namespace mlvm::IR {

std::string Function::debugString() const {
  std::stringstream ss{};
  ss << "Function: `" << name_ << "`\n";
  ss << "  Consts:\n";
  for (auto& c : consts_) {
    ss << "    - " << c->debugString() << "\n";
  }

  ss << "\n";
  ss << "  Instructions:\n";
  for (auto& ins : instructions_) {
    ss << "    - " << ins->debugString() << "\n";
  }

  ss << "\n";
  ss << "  Outputs:\n";
  for (auto& o : outputs_) {
    ss << "    - " << o->debugString() << "\n";
  }

  return ss.str();
};

}  // namespace mlvm::IR
