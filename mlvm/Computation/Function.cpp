#include "mlvm/Computation/Function.h"

#include <sstream>

#include "mlvm/Foundation/Macros.h"

namespace mlvm::computation {

using namespace array;
using namespace foundation;

StatusOr<TensorLike*> Function::makeTensor(ArrayLike arr) {
  MLVM_ASSIGN_OR_RETURN(arr_ptr, arr.get());

  auto next_id = constants_.size();
  std::stringstream name;
  name << "%c_" << next_id;

  auto tensor = new TensorLike{name.str(), std::move(arr_ptr), this};
  constants_.emplace_back(tensor);

  return constants_.back().get();
}

std::string Function::string() const {
  std::stringstream ss;
  ss << "Function \"" << name_ << "\" {\n";
  ss << "  Constants:\n";
  for (auto& c : constants_) {
    ss << "    " << c->string() << "\n";
  }
  ss << "}";
  return ss.str();
}

}  // namespace mlvm::computation
