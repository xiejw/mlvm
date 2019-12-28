#ifndef MLVM_ARRAY_ARRAY_
#define MLVM_ARRAY_ARRAY_

#include "mlvm/Array/Data.h"
#include "mlvm/Array/Shape.h"

#include "mlvm/Foundation/StatusOr.h"

namespace mlvm::array {

class Array {
 public:
  static foundation::StatusOr<Array> New(
      const std::initializer_list<double>& data,
      std::initializer_list<unsigned int> shape);

 public:
  std::string ToString() const;

 private:
  Array(Data data, Shape shape)
      : data_{std::move(data)}, shape_{std::move(shape)} {};

 private:
  Data data_;
  Shape shape_;
};

}  // namespace mlvm::array

#endif
