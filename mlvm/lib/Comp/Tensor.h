#ifndef MLVM_COMP_TENSOR
#define MLVM_COMP_TENSOR

namespace mlvm {
namespace comp {

class Tensor {
 public:
  enum class Kind { Constant };

  Kind kind() const { return kind_; }

 private:
  friend std::ostream& operator<<(std::ostream& os, const Tensor& arr) {
    return os;
  }

 private:
  Kind kind_;
};

}  // namespace comp
}  // namespace mlvm

#endif
