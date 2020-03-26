#ifndef MLVM_FOUNDATION_MACROS_H_
#define MLVM_FOUNDATION_MACROS_H_

#include <string>

#include "mlvm/Foundation/Utilities.h"

// Returns the function with error status if x.ok() == false.
#define MLVM_RETURN_IF_ERROR(x)      \
  {                                  \
    auto status = (x);               \
    if (!status.ok()) return status; \
  }

// Fatal error if x.ok() == false.
#define MLVM_FATAL_IF_ERROR(x)                                 \
  {                                                            \
    auto status = (x);                                         \
    if (!status.ok()) {                                        \
      mlvm::FatalError("Fatal error: %s",                      \
                       std::string(status.message()).c_str()); \
    }                                                          \
  }

#endif
