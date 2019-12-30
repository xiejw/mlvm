#include "mlvm/Computation/Function.h"

#include "mlvm/Array/Array.h"
#include "mlvm/Foundation/Macros.h"

namespace mlvm::computation {

using namespace foundation;

StatusOr<TensorLike*> Function::makeTensor(
    const std::initializer_list<double>& data,
    std::initializer_list<unsigned int> shape) {
  MLVM_ASSIGN_OR_RETURN(arr, array::Array::New(data, shape);
}

}  // namespace mlvm::computation
