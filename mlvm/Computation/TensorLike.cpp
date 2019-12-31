#include "mlvm/Computation/TensorLike.h"

#include <sstream>

namespace mlvm::computation {

std::string TensorLike::string() const {
  assert(type_ == Type::Constant);
  std::stringstream ss;
  ss << "`" << name_ << "`: C@";
  ss << array_->string();
  return ss.str();
}

const array::Shape& TensorLike::shape() const {
  if (type_ == Type::Constant) return array_->shape();

  assert(type_ == Type::Output);
  return output_shape_.value();
}

}  // namespace mlvm::computation
