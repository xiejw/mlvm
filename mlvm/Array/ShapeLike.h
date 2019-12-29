#ifndef MLVM_ARRAY_SHAPE_LIKE_
#define MLVM_ARRAY_SHAPE_LIKE_

#include <initializer_list>
#include <vector>

#include "mlvm/Array/Shape.h"
#include "mlvm/Foundation/StatusOr.h"

namespace mlvm::array {

// Represents a Shape Constructor.
class ShapeLike {
 public:
  ShapeLike(std::initializer_list<unsigned int> shape);

 public:
  foundation::StatusOr<Shape>&& Release() {
    return std::move(shape_or_.value());
  }

 private:
  std::optional<foundation::StatusOr<Shape>> shape_or_ = {};
};

}  // namespace mlvm::array

#endif
