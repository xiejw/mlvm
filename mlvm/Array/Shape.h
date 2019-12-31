#ifndef MLVM_ARRAY_SHAPE_
#define MLVM_ARRAY_SHAPE_

#include <initializer_list>
#include <string>
#include <vector>

namespace mlvm::array {

class ShapeLike;

// Represents a Shape.
//
// This should be super cheap to copy.
class Shape {
 public:
  Shape(const Shape&) = default;
  Shape& operator=(const Shape&) = default;
  Shape(Shape&&) = default;
  Shape& operator=(Shape&&) = default;

 public:
  // Debug string.
  std::string string() const;

  // Returns number of dimensions, i.e., Rank.
  //
  // For <3, 2>, rank is 2.
  unsigned int rank() const { return shape_.size(); };

  // Returns number of elements represented by this shape.
  //
  // For <3, 2>, element size is 3 * 2 = 6.
  unsigned int elementSize() const;

 private:
  friend class ShapeLike;

  Shape(const std::initializer_list<unsigned int>& shape) : shape_{shape} {};

 private:
  std::vector<unsigned int> shape_;
};

}  // namespace mlvm::array

#endif
