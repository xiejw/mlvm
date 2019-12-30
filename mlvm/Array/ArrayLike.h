#ifndef MLVM_ARRAY_ARRAY_LIKE_
#define MLVM_ARRAY_ARRAY_LIKE_

#include "mlvm/Array/Array.h"

#include "mlvm/Foundation/StatusOr.h"

#include "cassert"
#include "memory"
#include "optional"

namespace mlvm::array {

// A friendly builder for Array.
class ArrayLike {
  using StatusOrArray = foundation::StatusOr<std::unique_ptr<Array>>;

 public:
  ArrayLike(const std::initializer_list<double>& data,
            const std::initializer_list<unsigned int>& shape);

 public:
  StatusOrArray&& get() {
    assert(result_.has_value());
    return std::move(result_.value());
  };

 private:
  std::optional<StatusOrArray> result_ = {};
};

}  // namespace mlvm::array

#endif
