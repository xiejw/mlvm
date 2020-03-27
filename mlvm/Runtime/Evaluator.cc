#include "mlvm/Runtime/Evaluator.h"

#include <iostream>

using namespace mlvm::IR;

namespace mlvm::RT {

Status Evaluator::run(const Function& func) {
  std::cout << "Run Func: " << func.name();
  return Status::OK;
}

}  // namespace mlvm::RT
