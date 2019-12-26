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
  Shape(std::initializer_list<unsigned int> shape)
      : shape_{std::move(shape)} {

        };

 public:
  // Debug string.
  std::string DebugString() const;

  unsigned int Rank() const { return shape_.size(); };

 private:
  std::vector<unsigned int> shape_;
};

}  // namespace mlvm::array

#endif
