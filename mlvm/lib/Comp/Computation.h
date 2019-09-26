#ifndef MLVM_COMP_COMPUTATION
#define MLVM_COMP_COMPUTATION

#include <memory>
#include <string>
#include <vector>

#include "mlvm/lib/Comp/Instruction.h"
#include "mlvm/lib/Comp/OpType.h"
#include "mlvm/lib/Tensor/Tensor.h"

namespace mlvm {
namespace comp {

class Computation {
 public:
  const Instruction& newInstruction(
      std::string name, OpType op_type,
      std::initializer_list<tensor::Tensor> operands);

  std::string DebugString() const;

 private:
  std::vector<Instruction> ins_;
};

}  // namespace comp
}  // namespace mlvm

#endif
