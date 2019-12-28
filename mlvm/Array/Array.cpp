#include "mlvm/Array/Array.h"

#include <sstream>

namespace mlvm::array {

using namespace foundation;

StatusOr<Array> Array::New(const std::initializer_list<double>& data,
                           std::initializer_list<unsigned int> shape) {
  if (data.size() == 0)
    return Status::InvalidArguments("Data buffer for Array cannot be empty.");

  Data d{};
  d.Reset(data);

  auto shape_or = Shape::New(shape);
  if (!shape_or.Ok()) return shape_or.StatusOrDie();

  auto s = shape_or.ConsumeValue();
  if (s.ElementSize() != d.Size())
    return Status::InvalidArguments("Data and Shape sizes mismatch.");

  return Array{std::move(d), std::move(s)};
};

std::string Array::ToString() const {
  std::stringstream ss;
  ss << "[";
  ss << shape_.ToString();
  ss << " ";
  ss << data_.ToString();
  ss << "]";
  return ss.str();
}

}  // namespace mlvm::array
