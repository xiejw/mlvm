#ifndef MLVM_ARRAY_DATA
#define MLVM_ARRAY_DATA

#include <string>
#include <vector>

namespace mlvm {
namespace array {

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

}  // namespace array
}  // namespace mlvm

#endif
