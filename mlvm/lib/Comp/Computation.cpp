#include <iomanip>
#include <sstream>

#include "mlvm/lib/Comp/Computation.h"

namespace mlvm {
namespace comp {

const Instruction& Computation::newInstruction(std::string name) {
  return Instruction();
}

}  // namespace tensor
}  // namespace mlvm
