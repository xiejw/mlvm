#ifndef MLVM_FOUNDATION_STATUS_H_
#define MLVM_FOUNDATION_STATUS_H_

#include <optional>
#include <string>

#include "mlvm/Foundation/Macros.h"

namespace mlvm {

// Pre-defined error codes.
enum class ErrorCode {
  InvalidArguments,
  OSError,
  IOError,
};

class [[nodiscard]] Status {
 public:
  // Sets the error code. If present, allows an error message to be set.
  explicit Status(std::optional<ErrorCode> err,
                  std::optional<std::string> msg = std::optional<std::string>{})
      : err_{std::move(err)}, msg_{std::move(msg)} {};

  /// explicit Status(ErrorCode err) : err_{err} {};

  /// template <typename... T>
  /// explicit Status(ErrorCode err, const T&... args) : err_{err} {
  ///   msg_ = Strings::concat(args...);
  /// };

 public:
  // Pre-defined singleton.
  static const Status OK;

 public:
  // Returns true if no error.
  bool ok() const { return !err_.has_value(); }

  const ErrorCode& errorCode() const { return err_.value(); }
  std::string_view message() const {
    if (msg_.has_value()) return msg_.value();
    if (!err_.has_value()) return "(no error)";

    switch (err_.value()) {
      case ErrorCode::InvalidArguments:
        return "(unspecified invalid arguments error)";
      case ErrorCode::OSError:
        return "(unspecified os error)";
      case ErrorCode::IOError:
        return "(unspecified io error)";
      default:
        return "(unknown error)";
    }
  }

 private:
  std::optional<ErrorCode> err_;
  std::optional<std::string> msg_ = {};
};

}  // namespace mlvm

#endif
