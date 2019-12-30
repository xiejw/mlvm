#ifndef MLVM_ARRAY_SHAPE_LIKE_
#define MLVM_ARRAY_SHAPE_LIKE_

#include <cassert>
#include <initializer_list>
#include <vector>

#include "mlvm/Array/Shape.h"
#include "mlvm/Foundation/StatusOr.h"

namespace mlvm::array {

// A friendly constructor for Shape.
class ShapeLike {
  using StatusOrShape = foundation::StatusOr<Shape>;

 public:
  // Accepts initialization values.
  ShapeLike(const std::initializer_list<unsigned int>& shape);

  // Accepts a (moved) Shape.
  ShapeLike(Shape&& shape) : shape_or_{std::move(shape)} {};

 public:
  StatusOrShape&& get() {
    assert(shape_or_.has_value());
    return std::move(shape_or_.value());
  }

 private:
  std::optional<StatusOrShape> shape_or_ = {};
};

}  // namespace mlvm::array

#endif
