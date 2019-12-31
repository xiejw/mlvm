#include "mlvm/Array/ShapeLike.h"

namespace mlvm::array {

using namespace foundation;

ShapeLike::ShapeLike(const std::initializer_list<unsigned int>& shape) {
  if (shape.size() == 0) {
    shape_or_ = Status::InvalidArguments("Empty shape is not allowed.");
    return;
  }

  for (auto dim : shape) {
    if (dim <= 0) {
      shape_or_ = Status::InvalidArguments("Non-positive dim is not allowed.");
      return;
    }
  }
  shape_or_ = Shape(shape);
}

ShapeLike::ShapeLike(const Shape& shape) {
  Shape copy{shape};
  shape_or_ = std::move(copy);
}

}  // namespace mlvm::array
