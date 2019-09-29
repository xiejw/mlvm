#include <sstream>

#include "mlvm/lib/Array/Shape.h"

namespace mlvm {
namespace tensor {

std::string Shape::DebugString() const {
  std::stringstream ss;
  ss << *this;
  return ss.str();
}

std::ostream &operator<<(std::ostream &out, const Shape &s) {
  out << "<";

  int size = s.dims_->size();
  int i = 0;
  for (auto &d : *s.dims_) {
    if (++i != size)
      out << d << ", ";
    else
      out << d;
  }

  out << ">";
  return out;
}

}  // namespace tensor
}  // namespace mlvm
