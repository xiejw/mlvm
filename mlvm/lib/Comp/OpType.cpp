#include <iomanip>
#include <sstream>

#include "mlvm/lib/Comp/OpType.h"
#include "mlvm/lib/Support/Error.h"

namespace mlvm {
namespace comp {

std::ostream &operator<<(std::ostream &out, const OpType &s) {
  switch (s.kind_) {
    case OpType::OpAdd:
      out << "Add";
      break;
    default:
      mlvm::support::FatalError("Unknown support for OpType %d", int(s.kind_));
  }
  return out;
}

}  // namespace comp
}  // namespace mlvm
