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

// Returns the function with error status if status_or.ok() == false, where
// status_or = v.
//
// Otherwise, assigns the value to `x` (i.e., status_or.ok() == true).
// Here `x` should have type, consider to use
//
//    MLVM_ASSIGN_OR_RETURN(auto x, fn());
//
#define MLVM_ASSIGN_OR_RETURN(x, v) \
  MLVM_ASSIGN_OR_RETURN_IMPL(       \
      MLVM_STATUS_MACRO_CONCAT(__status_or, __COUNTER__), x, v)

// Similar to MLVM_ASSIGN_OR_RETURN, but triggers fatal error rather than
// return.
#define MLVM_ASSIGN_OR_FATAL(x, v) \
  MLVM_ASSIGN_OR_FATAL_IMPL(       \
      MLVM_STATUS_MACRO_CONCAT(__status_or, __COUNTER__), x, v)

///////////////////////////////////////////////////////////////////////////////
// Begin with helper macros.
///////////////////////////////////////////////////////////////////////////////

// For the reason why we need two macros here, check
// https://github.com/xiejw/eva/blob/master/docs/notes/iso.md#1933-the--operator-cppconcat
#define MLVM_STATUS_MACRO_CONCAT(x, y) MLVM_STATUS_MACRO_CONCAT_IMPL(x, y)
#define MLVM_STATUS_MACRO_CONCAT_IMPL(x, y) x##y

#define MLVM_ASSIGN_OR_RETURN_IMPL(so, x, y) \
  auto so = (y);                             \
  if (!so.ok()) return so.consumeStatus();   \
  x = so.consumeValue();

#define MLVM_ASSIGN_OR_FATAL_IMPL(so, x, y)                                  \
  auto so = (y);                                                             \
  if (!so.ok()) {                                                            \
    mlvm::FatalError("Fatal error: %s", so.statusOrDie().message().c_str()); \
  }                                                                          \
  x = so.consumeValue();

///////////////////////////////////////////////////////////////////////////////
// End with helper macros.
///////////////////////////////////////////////////////////////////////////////

#endif
