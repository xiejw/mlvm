#include <sstream>

#include "mlvm/lib/Comp/Computation.h"

namespace mlvm {
namespace comp {

const Instruction& Computation::newInstruction(
    std::string name, OpType op_type, std::initializer_list<Tensor> operands) {
  Instruction ins{name, op_type, operands};
  ins_.push_back(std::move(ins));
  return ins_.back();
}

std::string Computation::DebugString() const {
  std::stringstream ss;
  for (auto& ins : ins_) ss << ins << "\n";

  return ss.str();
}

}  // namespace comp
}  // namespace mlvm
