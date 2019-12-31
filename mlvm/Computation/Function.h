#ifndef MLVM_COMPUTATION_FUNCTION_
#define MLVM_COMPUTATION_FUNCTION_

#include "mlvm/Array/Array.h"
#include "mlvm/Computation/Instruction.h"
#include "mlvm/Computation/OpCode.h"
#include "mlvm/Computation/TensorLike.h"
#include "mlvm/Foundation/Macros.h"
#include "mlvm/Foundation/Status.h"
#include "mlvm/Foundation/StatusOr.h"

#include <string>
#include <vector>

namespace mlvm::computation {

class Function {
  using StatusOrPtrIns = foundation::StatusOr<Instruction*>;

 public:
  Function(std::string name) : name_{name} {};

 public:
  foundation::StatusOr<TensorLike*> makeTensor(array::ArrayLike arr);

  StatusOrPtrIns makeBinaryInst(OpCode op, TensorLike* const lhs,
                                TensorLike* const rhs) {
    return nullptr;
  };

  // Creates an Instruction grouping tensors as a Tuple.
  StatusOrPtrIns makeTupleInst(TensorLike* const lhs) { return nullptr; };

  // Sets the Output.
  StatusOrPtrIns setOutput(Instruction* const ins) { return nullptr; }

 public:
  // Debug string.
  std::string string() const;

 private:
  std::string name_;
  std::vector<std::unique_ptr<Instruction>> ins_vec_ = {};
  std::vector<std::unique_ptr<TensorLike>> constants_ = {};
};

}  // namespace mlvm::computation

#endif
