#ifndef MLVM_ARRAY_SHAPE_
#define MLVM_ARRAY_SHAPE_

#include <initializer_list>
#include <vector>

namespace mlvm::array {

// Represents a Shape.
//
// This should be super cheap to copy.
class Shape {
 public:
  Shape(std::initializer_list<unsigned int> shape) : shape_{std::move(shape)} {
    if (shape.size() == 0) throw "Empty shape is not allowed.";

    for (auto dim : shape) {
      if (dim <= 0) throw "Non-positive dim is not allowed.";
    }
  }

 public:
  // Debug string.
  std::string ToString() const;

  unsigned int Rank() const { return shape_.size(); };

 private:
  std::vector<unsigned int> shape_;
};

}  // namespace mlvm::array

#endif
