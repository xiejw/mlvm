#ifndef MLVM_COMP_COMPUTATION
#define MLVM_COMP_COMPUTATION

#include <memory>
#include <string>
#include <vector>

#include "mlvm/lib/Array/Array.h"
#include "mlvm/lib/Comp/Instruction.h"
#include "mlvm/lib/Comp/OpType.h"
#include "mlvm/lib/Comp/Tensor.h"

namespace mlvm {
namespace comp {

class Computation {
 public:
  const Instruction* newInstruction(std::string name, OpType op_type,
                                    std::initializer_list<Tensor*> operands);

  const Tensor* newConstant(std::string name, std::initializer_list<int> shape,
                            std::initializer_list<array::Float> data);

  std::string DebugString() const;

 private:
  std::vector<std::unique_ptr<Instruction>> ins_;
  std::vector<std::unique_ptr<Tensor>> tensors_;
};

}  // namespace comp
}  // namespace mlvm

#endif
