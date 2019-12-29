#include "mlvm/Array/Shape.h"

#include <sstream>

namespace mlvm::array {

unsigned int Shape::ElementSize() const {
  unsigned int size = 1;
  for (auto dim : shape_) size *= dim;
  return size;
}

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
