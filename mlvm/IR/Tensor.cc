#include "mlvm/IR/Tensor.h"

#include <absl/strings/str_cat.h>

namespace mlvm {

std::string Array::debugString() const {
  return absl::StrCat("`", value_, "`");
}

}  // namespace mlvm
