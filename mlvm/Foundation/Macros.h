#ifndef MLVM_FOUNDATION_MACROS_H_
#define MLVM_FOUNDATION_MACROS_H_

#include <cassert>
#include <string>

#include "mlvm/Foundation/Utilities.h"

#define MLVM_CHECK(x) assert(x)
#define MLVM_PRECONDITION(x)                             \
  {                                                      \
    if (!(x)) {                                          \
      FatalError("Precondition failed at %s", __FILE__); \
    }                                                    \
  }

// Returns the function with error status if x.ok() == false.
#define MLVM_RETURN_IF_ERROR(x)      \
  {                                  \
    auto status = (x);               \
    if (!status.ok()) return status; \
  }

// Fatal error if x.ok() == false.
#define MLVM_FATAL_IF_ERROR(x)                                       \
  {                                                                  \
    auto status = (x);                                               \
    if (!status.ok()) {                                              \
      mlvm::FatalError("Fatal error: %s", status.message().c_str()); \
    }                                                                \
  }

#endif
