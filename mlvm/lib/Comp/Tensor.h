#ifndef MLVM_COMP_TENSOR
#define MLVM_COMP_TENSOR

namespace mlvm {
namespace comp {

class Tensor {
 private:
  friend std::ostream& operator<<(std::ostream& os, const Tensor& arr) {
    return os;
  }
};

}  // namespace comp
}  // namespace mlvm

#endif
