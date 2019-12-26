#ifndef MLVM_ARRAY_DATA_
#define MLVM_ARRAY_DATA_

#include <cstddef>
#include <initializer_list>
#include <memory>
#include <string>

namespace mlvm::array {

// Represents a data buffer.
//
// - Empty buffer is invalid. It is interpreted as NotAllocated.
// - Buffer alias is not allowed.
class Data {
 public:
  // Debug string.
  std::string DebugString() const;

  // Check whether the data has been allocated.
  bool IsAllocated() const { return size_ > 0; }

  // Returns the number of elements.
  int Size() const { return size_; }

  // Move the `new_data` into this strucuture. So, owns the `new_data`.
  void Reset(double* new_data, std::size_t size);

  // Copy the `list` into this strucuture.
  void Reset(const std::initializer_list<double>& list);

 public:
  Data() {}
  Data(Data&&) = default;

  // Not allowed.
  Data(const Data&) = delete;
  Data& operator=(const Data&) = delete;

 private:
  std::unique_ptr<double[]> buf_;
  int size_;
};

}  // namespace mlvm::array

#endif
