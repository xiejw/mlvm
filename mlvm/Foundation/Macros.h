#ifndef MLVM_FOUNDATION_MACROS_H_
#define MLVM_FOUNDATION_MACROS_H_

// Returns the function with error status if x.ok() == false.
#define MLVM_RETURN_IF_ERROR(x)      \
  {                                  \
    auto status = (x);               \
    if (!status.ok()) return status; \
  }

#endif
