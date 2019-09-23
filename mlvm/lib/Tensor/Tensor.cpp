#include <sstream>

#include "mlvm/lib/Support/Error.h"
#include "mlvm/lib/Tensor/Tensor.h"

namespace mlvm {
namespace tensor {

const Data &Tensor::data() const {
  if (kind_ == Kind::Constant) {
    return *data_;
  }
  support::FatalError("Non constant Tensor does not have data.");
}

std::string Tensor::DebugString() const {
  std::stringstream ss;
  ss << *this;
  return ss.str();
}

std::ostream &operator<<(std::ostream &out, const Tensor &s) {
  out << "\"" << s.name_ << "\" ";
  out << s.shape_;
  return out;
}

}  // namespace tensor
}  // namespace mlvm
