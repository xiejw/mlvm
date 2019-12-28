#ifndef MLVM_ARRAY_SHAPE_
#define MLVM_ARRAY_SHAPE_

#include <initializer_list>
#include <vector>

#include "mlvm/Foundation/StatusOr.h"

namespace mlvm::array {

// Represents a Shape.
//
// This should be super cheap to copy.
class Shape {
 public:
  // Debug string.
  std::string ToString() const;

  unsigned int Rank() const { return shape_.size(); };

 private:
  Shape(std::initializer_list<unsigned int> shape)
      : shape_{std::move(shape)} {};

 public:
  static foundation::StatusOr<Shape> New(
      std::initializer_list<unsigned int> shape);

 private:
  std::vector<unsigned int> shape_;
};

}  // namespace mlvm::array

#endif
