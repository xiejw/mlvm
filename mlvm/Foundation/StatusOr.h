#ifndef MLVM_FOUNDATION_STATUSOR_H_
#define MLVM_FOUNDATION_STATUSOR_H_

#include <variant>

#include "mlvm/Foundation/Status.h"

namespace mlvm {

template <class T>
class StatusOr {
 public:
  StatusOr(T&& t) : value_{std::move(t)} {};

  StatusOr(Status&& status) : value_{std::move(status)} { };

  StatusOr(StatusOr&&) = default;
  StatusOr& operator=(StatusOr&&) = default;

  // Copy is not allowed.
  StatusOr(const StatusOr&) = delete;
  StatusOr& operator=(const StatusOr&) = delete;

 public:
  bool ok() const { return value_.index() == 1; };

  // // Requests ok() == false
  // const Status& statusOrDie() const {
  //   AssertNotHoldValue();
  //   return status_.value();
  // };

  // // Requests ok() == true
  // T& valueOrDie() {
  //   AssertHoldValue();
  //   AssertNotReleased();
  //   return value_.value();
  // };

  // // Requests ok() == false. Should be called at most once.
  // Status&& consumeStatus() {
  //   AssertNotHoldValue();
  //   return std::move(status_.value());
  // };

  // // Requests ok() == true. Should be called at most once.
  // T&& consumeValue() {
  //   AssertHoldValue();
  //   AssertNotReleased();
  //   Release();
  //   return std::move(value_.value());
  // };

 private:
  std::variant<Status, T> value_ = {};
};

}  // namespace mlvm

#endif
