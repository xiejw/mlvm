#include "mlvm/Computation/TensorLike.h"

#include <sstream>

#include "mlvm/Foundation/Macros.h"

namespace mlvm::computation {

std::string TensorLike::string() const {
  switch (type_) {
    case Type::Constant: {
      std::stringstream ss;
      ss << "`" << name_ << "`: C@";
      ss << array_->string();
      return ss.str();
    }
    case Type::Output: {
      std::stringstream ss;
      ss << "`" << name_ << "`: O@";
      ss << "[";
      ss << shape().string();
      ss << "]";
      return ss.str();
    }
    default:
      CHECK(false);
  }
}

const array::Shape& TensorLike::shape() const {
  if (type_ == Type::Constant) return array_->shape();

  CHECK(type_ == Type::Output);
  return output_shape_.value();
}

}  // namespace mlvm::computation
