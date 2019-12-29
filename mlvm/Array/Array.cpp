#include "mlvm/Array/Array.h"

#include <sstream>

namespace mlvm::array {

using namespace foundation;

StatusOr<Array> Array::New(const std::initializer_list<double>& data,
                           std::initializer_list<unsigned int> shape) {
  Data d{};
  MLVM_RETURN_IF_ERROR(d.reset(data));

  MLVM_ASSIGN_OR_RETURN(s, ShapeLike(shape).Get());

  if (s.elementSize() != d.size())
    return Status::InvalidArguments("Data and Shape sizes mismatch.");

  return Array{std::move(d), std::move(s)};
};

std::string Array::string() const {
  std::stringstream ss;
  ss << "[";
  ss << shape_.string();
  ss << " ";
  ss << data_.string();
  ss << "]";
  return ss.str();
}

}  // namespace mlvm::array
