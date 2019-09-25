#ifndef MLVM_COMP_COMPUTATION
#define MLVM_COMP_COMPUTATION

#include <memory>
#include <string>
#include <vector>

#include "mlvm/lib/Comp/Instruction.h"

namespace mlvm {
namespace comp {

class Computation {
 public:
  const Instruction& newInstruction(std::string name);

  std::string DebugString() const;

 private:
  std::vector<Instruction> ins_;
};

}  // namespace comp
}  // namespace mlvm

#endif
