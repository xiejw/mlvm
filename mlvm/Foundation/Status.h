#ifndef MLVM_FOUNDATION_STATUS_
#define MLVM_FOUNDATION_STATUS_

#include <cassert>
#include <memory>
#include <optional>
#include <string>

namespace mlvm::foundation {

#define MLVM_RETURN_IF_ERROR(x)      \
  {                                  \
    auto status = (x);               \
    if (!status.ok()) return status; \
  }

#define MLVM_STATUS_MACRO_CONCAT(x, y) x##y

#define MLVM_ASSIGN_OR_RETURN_IMPL(so, x, y) \
  auto so = (y);                             \
  if (!so.ok()) return so.consumeStatus();   \
  auto x = so.consumeValue();

#define MLVM_ASSIGN_OR_RETURN(x, y)                                            \
  MLVM_ASSIGN_OR_RETURN_IMPL(MLVM_STATUS_MACRO_CONCAT(status_or, __COUNTER__), \
                             x, y)

enum class ErrorCode {
  InvalidArguments,
};

class Status {
 public:
  // Sets the error code. If present, allows an error message to be set.
  explicit Status(std::optional<ErrorCode> err,
                  std::optional<std::string> msg = std::optional<std::string>{})
      : err_{std::move(err)}, msg_{std::move(msg)} {
    // `msg_` is present only if `err_` is present.
    assert(err_.has_value() || !msg_.has_value());
  };

  static const Status OK;

  static Status InvalidArguments(std::optional<std::string> msg = {}) {
    return Status(ErrorCode::InvalidArguments, msg);
  };

 public:
  // Returns true if no error.
  bool ok() const { return !err_.has_value(); }

  const ErrorCode& errorCode() const { return err_.value(); }
  const std::optional<std::string>& message() const { return msg_; }

 private:
  std::optional<ErrorCode> err_;
  std::optional<std::string> msg_;
};

}  // namespace mlvm::foundation

#endif
