#include "mlvm/Foundation/Logging.h"

#include "absl/flags/flag.h"

ABSL_FLAG(int, v, 0, "Logging level (defaults to `0`)");

namespace mlvm::logging {

VoidType VoidType::instance = VoidType{};

}  // namespace mlvm::logging
