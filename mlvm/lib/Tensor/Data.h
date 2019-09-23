#ifndef MLVM_TENSOR_DATA
#define MLVM_TENSOR_DATA

#include <string>
#include <vector>

namespace mlvm {
namespace tensor {

using Float = float;

// Data structure holding the Float data.
class Data {
 public:
  explicit Data(std::initializer_list<Float> data) : data_{data} {}

  std::string DebugString() const;

  // const std::vector<int>& dims() const { return *dims_; };

 private:
  friend std::ostream& operator<<(std::ostream& os, const Data& s);

 private:
  std::vector<Float> data_;
};

}  // namespace tensor
}  // namespace mlvm

#endif
