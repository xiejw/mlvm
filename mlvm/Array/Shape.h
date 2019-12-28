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

  // Returns number of dimensions, i.e., Rank.
  //
  // For <3, 2>, rank is 2.
  unsigned int Rank() const { return shape_.size(); };


  // Returns number of elements represented by this shape.
  //
  // For <3, 2>, element size is 3 * 2 = 6.
  unsigned int ElementSize() const;

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
