#include <sstream>

#include "mlvm/lib/Comp/Computation.h"

namespace mlvm {
namespace comp {

const Instruction* Computation::newInstruction(
    std::string name, OpType op_type, std::initializer_list<Tensor*> operands) {
  std::unique_ptr<Instruction> ins{new Instruction{name, op_type, operands}};
  ins_.push_back(std::move(ins));
  return ins_.back().get();
}

const Tensor* Computation::newConstant(
    std::string name, std::initializer_list<int> shape,
    std::initializer_list<array::Float> data) {
  std::unique_ptr<array::Array> arr{new array::Array(name, shape, data)};
  std::unique_ptr<Tensor> t{new Tensor(std::move(arr))};
  tensors_.push_back(std::move(t));
  return tensors_.back().get();
}

std::string Computation::DebugString() const {
  std::stringstream ss;
  for (auto& ins : ins_) ss << *ins << "\n";

  return ss.str();
}

}  // namespace comp
}  // namespace mlvm
