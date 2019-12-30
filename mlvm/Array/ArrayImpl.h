#ifndef MLVM_ARRAY_ARRAY_
#define MLVM_ARRAY_ARRAY_

#include "mlvm/Array/Data.h"
#include "mlvm/Array/Shape.h"

#include "memory"

namespace mlvm::array {

class Array {
 public:
  Array(Array&&) = default;
  Array& operator=(Array&& other) = default;

 public:
  std::string string() const;

 private:
  friend class ArrayLike;

  Array(Data data, Shape shape)
      : data_{std::move(data)}, shape_{std::move(shape)} {};

 private:
  Data data_;
  Shape shape_;
};

}  // namespace mlvm::array

#endif
