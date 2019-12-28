#include "mlvm/Array/Array.h"

#include <sstream>

namespace mlvm::array {

using namespace foundation;

StatusOr<Array> Array::New(const std::initializer_list<double>& data,
                           std::initializer_list<unsigned int> shape) {
  if (data.size() == 0)
    return Status::InvalidArguments("Data buffer for Array cannot be empty.");

  // element match.
  Data d{};
  d.Reset(data);

  auto shape_or = Shape::New(shape);
  if (!shape_or.Ok()) return shape_or.StatusOrDie();

  return Array{std::move(d), shape_or.ConsumeValue()};
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
