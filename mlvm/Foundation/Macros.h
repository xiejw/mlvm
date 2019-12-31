#ifndef MLVM_FOUNDATION_MACROS_
#define MLVM_FOUNDATION_MACROS_

#define MLVM_RETURN_IF_ERROR(x)      \
  {                                  \
    auto status = (x);               \
    if (!status.ok()) return status; \
  }

#define MLVM_STATUS_MACRO_CONCAT(x, y) MLVM_STATUS_MACRO_CONCAT_IMPL(x, y)
#define MLVM_STATUS_MACRO_CONCAT_IMPL(x, y) x##y

#define MLVM_ASSIGN_OR_RETURN(x, y)                                            \
  MLVM_ASSIGN_OR_RETURN_IMPL(MLVM_STATUS_MACRO_CONCAT(status_or, __COUNTER__), \
                             x, y)

#define MLVM_ASSIGN_OR_RETURN_IMPL(so, x, y) \
  auto so = (y);                             \
  if (!so.ok()) return so.consumeStatus();   \
  auto x = so.consumeValue();


#endif
