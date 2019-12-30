#ifndef MLVM_COMPUTATION_FUNCTION_
#define MLVM_COMPUTATION_FUNCTION_

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
 public:
  Function(std::string name) : ins_vec_{}, name_{name} {};

 public:
  foundation::StatusOr<TensorLike*> makeTensor(
      const std::initializer_list<double>& data,
      std::initializer_list<unsigned int> shape);

  foundation::StatusOr<Instruction*> addBinaryInst(OpCode op,
                                                   TensorLike* const lhs,
                                                   TensorLike* const rhs) {
    return nullptr;
  };

  foundation::StatusOr<Instruction*> addTupleInst(TensorLike* const lhs) {
    return nullptr;
  };

  foundation::StatusOr<Instruction*> setOutput(Instruction* const ins) {
    return nullptr;
  }

 private:
  std::vector<std::unique_ptr<Instruction>> ins_vec_;
  std::string name_;
};

}  // namespace mlvm::computation

#endif
