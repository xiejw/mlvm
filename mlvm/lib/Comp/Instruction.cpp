#include <iomanip>
#include <sstream>

#include "mlvm/lib/Comp/Instruction.h"

namespace mlvm {
namespace comp {

std::string Instruction::DebugString() const {
  std::stringstream ss;

  ss << "\"" << name << "\" (" << op_type << ")";

  return ss.str();
}

}  // namespace comp
}  // namespace mlvm
