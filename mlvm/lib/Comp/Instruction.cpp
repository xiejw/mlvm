#include <sstream>

#include "mlvm/lib/Comp/Instruction.h"

namespace mlvm {
namespace comp {

std::string Instruction::DebugString() const {
  std::stringstream ss;
  ss << *this;
  return ss.str();
}

std::ostream &operator<<(std::ostream &out, const Instruction &s) {
  out << "\"" << s.name << "\" (" << s.op_type << ")";
  return out;
}

}  // namespace comp
}  // namespace mlvm
