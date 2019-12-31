#ifndef MLVM_ARRAY_ARRAY_LIKE_
#define MLVM_ARRAY_ARRAY_LIKE_

#include "mlvm/Array/ArrayImpl.h"

#include "mlvm/Foundation/Macros.h"
#include "mlvm/Foundation/StatusOr.h"

#include "memory"
#include "optional"

namespace mlvm::array {

// A friendly constructor for Array.
class ArrayLike {
  using StatusOrPtrArray = foundation::StatusOr<std::unique_ptr<Array>>;

 public:
  // Accepts initialization values.
  ArrayLike(const std::initializer_list<double>& data,
            const std::initializer_list<unsigned int>& shape);

  // Accepts an Array in heap.
  ArrayLike(std::unique_ptr<Array> arr);

 public:
  StatusOrPtrArray&& get() {
    CHECK(array_or_.has_value());
    return std::move(array_or_.value());
  };

 private:
  std::optional<StatusOrPtrArray> array_or_ = {};
};

}  // namespace mlvm::array

#endif
