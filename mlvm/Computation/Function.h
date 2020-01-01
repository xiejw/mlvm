#ifndef MLVM_COMPUTATION_FUNCTION_
#define MLVM_COMPUTATION_FUNCTION_

#include "mlvm/Array/Array.h"
#include "mlvm/Computation/Instruction.h"
#include "mlvm/Computation/OpCode.h"
#include "mlvm/Computation/Tuple.h"
#include "mlvm/Computation/TensorLike.h"
#include "mlvm/Foundation/Macros.h"
#include "mlvm/Foundation/Status.h"
#include "mlvm/Foundation/StatusOr.h"

#include <string>
#include <vector>

namespace mlvm::computation {

class Function {
  using StatusOrPtrIns = foundation::StatusOr<Instruction*>;
  using StatusOrPtrTensor = foundation::StatusOr<TensorLike*>;

 public:
  Function(std::string name) : name_{name} {};

 public:
  StatusOrPtrTensor makeTensor(array::ArrayLike arr);

  StatusOrPtrIns makeBinaryInst(OpCode op, TensorLike* const lhs,
                                TensorLike* const rhs);

  // Sets the Output.
  foundation::Status setOutput(std::unique_ptr<Tuple> output) {
    // Output cannot be Placeholder.
    output_.swap(output);
    return foundation::Status::OK;
  }

 public:
  // Debug string.
  std::string string() const;

  const Tuple* output() const { return output_.get(); }

 private:
  std::string name_;
  std::vector<std::unique_ptr<Instruction>> ins_vec_ = {};
  std::vector<std::unique_ptr<TensorLike>> constants_ = {};

  std::unique_ptr<Tuple> output_ = {};
};

}  // namespace mlvm::computation

#endif
