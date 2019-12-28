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

  const Status& StatusOrDie() const { return status_.value(); };
  const T& ValueOrDie() const { return value_.value(); };

  T&& ConsumeValue() { return std::move(value_.value()); };

 private:
  std::optional<T> value_ = {};
  std::optional<Status> status_ = {};
};

}  // namespace mlvm::foundation

#endif
