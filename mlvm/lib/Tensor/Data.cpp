#include <iomanip>
#include <sstream>

#include "mlvm/lib/Tensor/Data.h"

namespace mlvm {
namespace tensor {

std::string Data::DebugString() const {
  std::stringstream ss;
  ss << *this;
  return ss.str();
}

std::ostream &operator<<(std::ostream &out, const Data &s) {
  out << std::fixed << std::setprecision(3) << "[";

  int size = s.data_.size();
  int i = 0;
  for (auto &d : s.data_) {
    if (++i != size)
      out << d << ", ";
    else
      out << d;
  }

  out << "]";
  return out;
}

}  // namespace tensor
}  // namespace mlvm
