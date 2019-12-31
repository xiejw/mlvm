#ifndef MLVM_COMPUTATION_TENSORLIKE_
#define MLVM_COMPUTATION_TENSORLIKE_

#include <memory>

#include "mlvm/Array/Array.h"

namespace mlvm::computation {

class TensorLike {
 public:
  enum class Type {
    Array,
  };

 public:
  Type type() const { return type_; }

 private:
  friend class Function;

  TensorLike(std::unique_ptr<array::Array> arr) : type_{Type::Array} {
    array_.swap(arr);
  };

 private:
  Type type_;
  std::unique_ptr<array::Array> array_ = {};
};

}  // namespace mlvm::computation

#endif
