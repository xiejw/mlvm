#ifndef MLVM_FOUNDATION_STATUSOR_H_
#define MLVM_FOUNDATION_STATUSOR_H_

#include <variant>

#include "mlvm/Foundation/Status.h"
#include "mlvm/Foundation/Utilities.h"

namespace mlvm {

template <class T>
class StatusOr {
 public:
  StatusOr(T&& t) : value_{std::move(t)} {};

  StatusOr(Status&& status) : value_{std::move(status)} {
    MLVM_PRECONDITION(!status.ok());
  };

  StatusOr(StatusOr&&) = default;
  StatusOr& operator=(StatusOr&&) = default;

  // Copy is not allowed.
  StatusOr(const StatusOr&) = delete;
  StatusOr& operator=(const StatusOr&) = delete;

 public:
  bool ok() const { return value_.index() == 1; };

  // Requests ok() == false
  const Status& statusOrDie() const {
    MLVM_PRECONDITION(!ok());
    return std::get<0>(value_);
  };

  // Requests ok() == false. Should be called at most once.
  Status&& consumeStatus() {
    MLVM_PRECONDITION(!ok());
    return std::move(std::get<0>(value_));
  };

  // Requests ok() == true
  T& valueOrDie() {
    MLVM_PRECONDITION(ok());
    return std::get<1>(value_);
  };

  // Requests ok() == true. Should be called at most once.
  T&& consumeValue() {
    MLVM_PRECONDITION(ok());
    return std::move(std::get<1>(value_));
  };

 private:
  std::variant<Status, T> value_ = {};
};

}  // namespace mlvm

#endif
