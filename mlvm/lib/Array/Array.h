#ifndef MLVM_ARRAY_ARRAY
#define MLVM_ARRAY_ARRAY

#include <memory>
#include <string>
#include <vector>

#include "mlvm/lib/Array/Data.h"
#include "mlvm/lib/Array/Shape.h"

namespace mlvm {
namespace array {

// Immutable structure holding the Constant Tensor information.
class Array {
 public:
  explicit Array(std::string name, std::initializer_list<int> shape,
                 std::initializer_list<Float> data)
      : name_{name}, shape_{shape}, data_{data} {}

  std::string DebugString() const;

  const Data& data() const;

  const Shape& shape() const;

 private:
  friend std::ostream& operator<<(std::ostream& os, const Array& arr);

 private:
  std::string name_;
  Shape shape_;
  Data data_;
};

}  // namespace array
}  // namespace mlvm

#endif
