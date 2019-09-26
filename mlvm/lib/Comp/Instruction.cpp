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
  // Name and Op Type.
  out << "\"" << s.name << "\" (" << s.op_type << ")";

  // Operands
  out << " (";
  for (auto &o: s.operands) {
    out << o;
  }
  out << ") -> ";

  // Results
  out << "()";
  return out;
}

}  // namespace comp
}  // namespace mlvm
