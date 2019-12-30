#include "mlvm/Computation/Function.h"

#include "mlvm/Foundation/Macros.h"

namespace mlvm::computation {

using namespace array;
using namespace foundation;

StatusOr<TensorLike*> Function::makeTensor(ArrayLike arr) {
  MLVM_ASSIGN_OR_RETURN(arr_ptr, arr.get());
}

}  // namespace mlvm::computation
