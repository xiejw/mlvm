#ifndef MLVM_FOUNDATION_STATUSOR_
#define MLVM_FOUNDATION_STATUSOR_

#include <cassert>
#include <optional>

#include "mlvm/Foundation/Status.h"

namespace mlvm::foundation {

template <class T>
class StatusOr {
 public:
  explicit StatusOr(T t) : value_{std::move(t)} {};

  explicit StatusOr(Status status) : status_{std::move(status)} {
    assert(status_.has_value());
    assert(!status_.value().Ok());
  };

 public:
  bool Ok() const { return value_.has_value(); };

  // Requests Ok() == false
  const Status& StatusOrDie() const { return status_.value(); };
  // Requests Ok() == true
  const T& ValueOrDie() const {
    AssertNotReleased();
    return value_.value();
  };

  // Requests Ok() == true. Should be called at most once.
  T&& ConsumeValue() {
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
  void AssertNotReleased() const { assert(!released_); };
  void Release() { released_ = true; }
#else
  void inline AssertNotReleased() const {};
  void inline Release(){};
#endif
};

}  // namespace mlvm::foundation

#endif
