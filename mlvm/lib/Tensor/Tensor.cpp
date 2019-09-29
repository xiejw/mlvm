#include <sstream>

#include "mlvm/lib/Tensor/Tensor.h"

namespace mlvm {
namespace tensor {

const Data &Array::data() const { return data_; }

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

}  // namespace tensor
}  // namespace mlvm
