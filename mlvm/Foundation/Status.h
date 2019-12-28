#ifndef MLVM_FOUNDATION_STATUS_
#define MLVM_FOUNDATION_STATUS_

#include <cassert>
#include <memory>
#include <optional>
#include <string>

namespace mlvm::foundation {

enum class ErrorCode {
  INVALID_ARGUMENTS,
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
    return Status(ErrorCode::INVALID_ARGUMENTS, msg);
  };

 public:
  // Returns true if no error.
  bool Ok() const { return !err_.has_value(); }

  const ErrorCode& Error() const { return err_.value(); }
  const std::optional<std::string>& Message() const { return msg_; }

 private:
  std::optional<ErrorCode> err_;
  std::optional<std::string> msg_;
};

}  // namespace mlvm::foundation

#endif
