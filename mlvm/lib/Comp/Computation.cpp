#include <iomanip>
#include <sstream>

#include "mlvm/lib/Comp/Computation.h"

namespace mlvm {
namespace comp {

const Instruction& Computation::newInstruction(std::string name) {
  Instruction ins{name};
  ins_.push_back(std::move(ins));
  return ins_.back();
}

std::string Computation::DebugString() const {
  std::stringstream ss;
  for (auto& ins : ins_) ss << ins.name << "\n";

  return ss.str();
}

}  // namespace comp
}  // namespace mlvm
