#include "mlvm/IR/Tensor.h"

#include <absl/strings/str_cat.h>

namespace mlvm {

std::string Array::debugString() const {
  return absl::StrCat("`", value_, "`");
}

std::string OutputTensor::debugString() const {
  return absl::StrCat("`", name_, "`");
}

}  // namespace mlvm
