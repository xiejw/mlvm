#include "mlvm/Computation/TensorLike.h"

#include <sstream>

#include "mlvm/Foundation/Macros.h"

namespace mlvm::computation {

std::string TensorLike::string() const {
  CHECK(type_ == Type::Constant);
  std::stringstream ss;
  ss << "`" << name_ << "`: C@";
  ss << array_->string();
  return ss.str();
}

const array::Shape& TensorLike::shape() const {
  if (type_ == Type::Constant) return array_->shape();

  CHECK(type_ == Type::Output);
  return output_shape_.value();
}

}  // namespace mlvm::computation
