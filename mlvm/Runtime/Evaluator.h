#ifndef MLVM_RUNTIME_EVALUATOR_H_
#define MLVM_RUNTIME_EVALUATOR_H_

#include <string>

#include "mlvm/Foundation/StatusOr.h"
#include "mlvm/IR/Function.h"

namespace mlvm::RT {

class Evaluator {
 public:
  Status run(const IR::Function& func);
};

}  // namespace mlvm::RT

#endif
