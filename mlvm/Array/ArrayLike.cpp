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
      result_ = std::move(status);
      return;
    }
  }

  auto shape_or = ShapeLike(shape).Get();
  if (!shape_or.ok()) {
    result_ = shape_or.consumeStatus();
    return;
  }
  auto s = shape_or.consumeValue();

  if (s.elementSize() != d.size()) {
    result_ = Status::InvalidArguments("Data and Shape sizes mismatch.");
    return;
  }

  result_ = std::unique_ptr<Array>{new Array{std::move(d), std::move(s)}};
};

}  // namespace mlvm::array
