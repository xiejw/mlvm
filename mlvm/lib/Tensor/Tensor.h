#ifndef MLVM_TENSOR_TENSOR
#define MLVM_TENSOR_TENSOR

#include <memory>
#include <string>
#include <vector>

#include "mlvm/lib/Tensor/Data.h"
#include "mlvm/lib/Tensor/Shape.h"

namespace mlvm {
namespace tensor {

// Immutable structure holding the Tensor information.
//
// Copy is recommended to share the Tensor.
class Tensor {
 public:
  enum Kind { Constant };

  Tensor() = delete;

  std::string DebugString() const;

 public:
  static Tensor newConstant(std::string name, std::initializer_list<int> shape,
                            std::initializer_list<Float> data) {
    return Tensor(Kind::Constant, name, shape, data);
  }

  const Data& data() const;

 private:
  explicit Tensor(Kind kind, std::string name, std::initializer_list<int> shape,
                  std::initializer_list<Float> data)
      : kind_{kind}, name_{name}, shape_{shape}, data_{new Data{data}} {}

  friend std::ostream& operator<<(std::ostream& os, const Tensor& s);

 private:
  Kind kind_;
  std::string name_;
  Shape shape_;
  std::shared_ptr<Data> data_;
};

}  // namespace tensor
}  // namespace mlvm

#endif
