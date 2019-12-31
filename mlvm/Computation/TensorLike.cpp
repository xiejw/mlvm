#include "mlvm/Computation/TensorLike.h"

#include <sstream>

namespace mlvm::computation {

std::string TensorLike::string() const {
  assert(type_ == Type::Array);
  std::stringstream ss;
  ss << "`" << name_ << "`: C@";
  ss << array_->string();
  return ss.str();
}

}  // namespace mlvm::computation
