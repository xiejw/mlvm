#ifndef MLVM_COMPUTATION_TUPLE_
#define MLVM_COMPUTATION_TUPLE_

#include <vector>

#include "mlvm/Computation/TensorLike.h"

namespace mlvm::computation {

// Immutable structure for inputs and outputs.
class Tuple {
 public:
  Tuple(std::vector<TensorLike*> items) : items_{std::move(items)} {};

  const std::vector<TensorLike*> items() const { return items_; }

 private:
  std::vector<TensorLike*> items_;
};

}  // namespace mlvm::computation

#endif
