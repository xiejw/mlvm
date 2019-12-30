#ifndef MLVM_ARRAY_SHAPE_LIKE_
#define MLVM_ARRAY_SHAPE_LIKE_

#include <cassert>
#include <initializer_list>
#include <vector>

#include "mlvm/Array/Shape.h"
#include "mlvm/Foundation/StatusOr.h"

namespace mlvm::array {

// Represents a Shape Constructor.
//
// This makes constructing Shape much easier.
class ShapeLike {
 public:
  ShapeLike(const std::initializer_list<unsigned int>& shape);
  ShapeLike(Shape&& shape) : shape_or_{std::move(shape)} {};

 public:
  foundation::StatusOr<Shape>&& Get() {
    assert(shape_or_.has_value());
    return std::move(shape_or_.value());
  }

  Shape&& ShapeOrDie() {
    assert(shape_or_.has_value() && shape_or_.value().ok());
    return shape_or_.value().consumeValue();
  }

 private:
  std::optional<foundation::StatusOr<Shape>> shape_or_ = {};
};

}  // namespace mlvm::array

#endif
