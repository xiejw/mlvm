#include <sstream>

#include "mlvm/lib/Array/Array.h"

namespace mlvm {
namespace array {

const Data &Array::data() const { return data_; }

const Shape &Array::shape() const { return shape_; }

std::string Array::DebugString() const {
  std::stringstream ss;
  ss << *this;
  return ss.str();
}

std::ostream &operator<<(std::ostream &out, const Array &arr) {
  out << "\"" << arr.name_ << "\" ";
  out << arr.shape_;
  return out;
}

}  // namespace array
}  // namespace mlvm
