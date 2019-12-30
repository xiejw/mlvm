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
  foundation::StatusOr<Instruction*> addBinaryInst(OpCode op,
                                                   const TensorLike& lhs,
                                                   const TensorLike& rhs) {
    return nullptr;
  };

  foundation::StatusOr<Instruction*> addTupleInst(const TensorLike& lhs) {
    return nullptr;
  };

  foundation::StatusOr<Instruction*> setOutput(Instruction* const ins) {
    return nullptr;
  }

 private:
  std::vector<std::unique_ptr<Instruction>> ins_vec_;
  ;
  std::string name_;
};

}  // namespace mlvm::computation

#endif
