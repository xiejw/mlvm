#include "mlvm/Array/ShapeLike.h"

namespace mlvm::array {

using namespace foundation;

ShapeLike::ShapeLike(std::initializer_list<unsigned int> shape) {
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
  shape_or_ = Shape(std::move(shape));
}

}  // namespace mlvm::array
