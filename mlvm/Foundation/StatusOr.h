#ifndef MLVM_FOUNDATION_STATUSOR_
#define MLVM_FOUNDATION_STATUSOR_

#include <optional>

#include "mlvm/Foundation/Macros.h"
#include "mlvm/Foundation/Status.h"

namespace mlvm::foundation {

template <class T>
class StatusOr {
 public:
  StatusOr(T&& t) : value_{std::move(t)} {};

  StatusOr(Status&& status) : status_{std::move(status)} {
    CHECK(status_.has_value());
    CHECK(!status_.value().ok());
  };

  StatusOr(StatusOr&&) = default;
  StatusOr& operator=(StatusOr&&) = default;

  // Copy is now allowed.
  StatusOr(const StatusOr&) = delete;
  StatusOr& operator=(const StatusOr&) = delete;

 public:
  bool ok() const { return value_.has_value(); };

  // Requests ok() == false
  const Status& statusOrDie() const {
    AssertNotHoldValue();
    return status_.value();
  };

  // Requests ok() == true
  T& valueOrDie() {
    AssertHoldValue();
    AssertNotReleased();
    return value_.value();
  };

  // Requests ok() == false. Should be called at most once.
  Status&& consumeStatus() {
    AssertNotHoldValue();
    return std::move(status_.value());
  };

  // Requests ok() == true. Should be called at most once.
  T&& consumeValue() {
    AssertHoldValue();
    AssertNotReleased();
    Release();
    return std::move(value_.value());
  };

 private:
  std::optional<T> value_ = {};
  std::optional<Status> status_ = {};

 private:
#ifndef NDEBUG
  bool released_ = false;
  void inline AssertNotReleased() const { CHECK(!released_); };
  void inline Release() { released_ = true; }
  void inline AssertHoldValue() const { CHECK(value_.has_value()); };
  void inline AssertNotHoldValue() const { CHECK(!value_.has_value()); };
#else
  void inline AssertNotReleased() const {};
  void inline Release(){};
  void inline AssertHoldValue() const {};
  void inline AssertNotHoldValue() const {};
#endif
};

}  // namespace mlvm::foundation

#endif
