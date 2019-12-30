#include "mlvm/Array/ArrayLike.h"

#include "mlvm/Array/ShapeLike.h"

namespace mlvm::array {

using namespace foundation;

ArrayLike::ArrayLike(const std::initializer_list<double>& data,
                     const std::initializer_list<unsigned int>& shape) {
  Data d{};
  {
    auto status = d.reset(data);
    if (!status.ok()) {
      array_or_ = std::move(status);
      return;
    }
  }

  auto shape_or = ShapeLike(shape).get();
  if (!shape_or.ok()) {
    array_or_ = shape_or.consumeStatus();
    return;
  }
  auto s = shape_or.consumeValue();

  if (s.elementSize() != d.size()) {
    array_or_ = Status::InvalidArguments("Data and Shape sizes mismatch.");
    return;
  }

  array_or_ = std::unique_ptr<Array>{new Array{std::move(d), std::move(s)}};
};

ArrayLike::ArrayLike(std::unique_ptr<Array> arr) {
  array_or_ = std::move(arr);
};

}  // namespace mlvm::array
