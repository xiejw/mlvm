#include "mlvm/Array/ArrayImpl.h"

#include <sstream>

namespace mlvm::array {

using namespace foundation;

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
