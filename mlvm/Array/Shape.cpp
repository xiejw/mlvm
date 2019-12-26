#include "mlvm/Array/Shape.h"

#include <sstream>

namespace mlvm::array {

std::string Shape::ToString() const {
  std::stringstream ss;
  ss << "<";
  int size = shape_.size();
  for (int i = 0; i < size; i++) {
    ss << shape_[i];
    if (i != size - 1) ss << ", ";
  }
  ss << ">";
  return ss.str();
}

}  // namespace mlvm::array
