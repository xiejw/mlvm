#include "mlvm/Computation/Function.h"

#include <sstream>

#include "mlvm/Foundation/Macros.h"

namespace mlvm::computation {

using namespace array;
using namespace foundation;

StatusOr<TensorLike*> Function::makeTensor(ArrayLike arr) {
  MLVM_ASSIGN_OR_RETURN(arr_ptr, arr.get());

  auto tensor = new TensorLike{std::move(arr_ptr)};
  constants_.emplace_back(tensor);

  return Status::InvalidArguments("1");
}

std::string Function::string() const {
  std::stringstream ss;
  ss << "Function \"" << name_ << "\" {\n";
  ss << "}";
  return ss.str();
}

}  // namespace mlvm::computation
