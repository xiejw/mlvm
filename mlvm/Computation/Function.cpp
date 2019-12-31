#include "mlvm/Computation/Function.h"

#include <sstream>

#include "mlvm/Foundation/Macros.h"

namespace mlvm::computation {

using namespace array;
using namespace foundation;

Function::StatusOrPtrTensor Function::makeTensor(ArrayLike arr) {
  MLVM_ASSIGN_OR_RETURN(arr_ptr, arr.get());

  // Constant name.
  auto next_id = constants_.size();
  std::stringstream name;
  name << "%c_" << next_id;

  auto tensor = new TensorLike{name.str(), std::move(arr_ptr), this};
  constants_.emplace_back(tensor);

  return constants_.back().get();
}

Function::StatusOrPtrIns Function::makeBinaryInst(OpCode op,
                                                  TensorLike* const lhs,
                                                  TensorLike* const rhs) {
  assert(op == OpCode::Add);

  // Instruction name.
  auto next_id = ins_vec_.size();
  std::stringstream name;
  name << "%" << next_id;

  auto ins = new Instruction {name.str(), op, {lhs, rhs}};
  ins_vec_.emplace_back(ins);
  return ins_vec_.back().get();
}

std::string Function::string() const {
  std::stringstream ss;
  ss << "Function \"" << name_ << "\" {\n";

  ss << "  Constants:\n";
  for (auto& c : constants_) {
    ss << "    " << c->string() << "\n";
  }

  ss << "\n  Instructions:\n";
  for (auto& ins : ins_vec_) {
    ss << "    " << ins->string() << "\n";
  }
  ss << "}";
  return ss.str();
}

}  // namespace mlvm::computation
