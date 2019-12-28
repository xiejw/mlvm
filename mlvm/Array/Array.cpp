#include "mlvm/Array/Array.h"

#include <sstream>

namespace mlvm::array {

using namespace foundation;

StatusOr<Array> Array::New(const std::initializer_list<double>& data,
                           std::initializer_list<unsigned int> shape) {
  Data d{};
  MLVM_RETURN_IF_ERROR(d.Reset(data));

  MLVM_ASSIGN_OR_RETURN(s, Shape::New(shape));

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
