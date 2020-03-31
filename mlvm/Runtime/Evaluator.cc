#include "mlvm/Runtime/Evaluator.h"

#include "mlvm/Foundation/Logging.h"

#include <iostream>

using namespace mlvm::IR;

namespace mlvm::RT {

Status Evaluator::run(const Function& func) {
  LOG_INFO() << "Run Func: " << func.name();
  LOG_DEBUG() << "Func:\n" << func.debugString();
  return Status::OK;
}

}  // namespace mlvm::RT
