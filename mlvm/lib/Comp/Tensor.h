#ifndef MLVM_COMP_TENSOR
#define MLVM_COMP_TENSOR

#include <memory>
#include <string>

#include "mlvm/lib/Array/Array.h"

namespace mlvm {
namespace comp {

class Tensor {
 public:
  enum class Kind { Constant };

  Kind kind() const { return kind_; }

 private:
  explicit Tensor(Kind kind, std::unique_ptr<array::Array> arr)
      : kind_{kind}, array_{arr.release()} {}

  friend std::ostream& operator<<(std::ostream& os, const Tensor& arr) {
    return os;
  }

 private:
  Kind kind_;
  std::unique_ptr<array::Array> array_;
};

}  // namespace comp
}  // namespace mlvm

#endif
