#ifndef MLVM_ARRAY_DATA_
#define MLVM_ARRAY_DATA_

#include <cstddef>
#include <initializer_list>
#include <memory>
#include <string>

#include "mlvm/Foundation/Status.h"

namespace mlvm::array {

// Represents a data buffer.
//
// - Empty buffer is invalid. It is interpreted as NotAllocated.
// - Buffer alias is not allowed.
class Data {
 public:
  // Debug string.
  std::string ToString() const;

  // Check whether the data has been allocated.
  bool IsAllocated() const { return size_ > 0; }

  // Returns the number of elements.
  int Size() const { return size_; }

  // Move the `new_data` into this strucuture. So, owns the `new_data`.
  foundation::Status Reset(double* new_data, std::size_t size);

  // Copy the `list` into this strucuture.
  foundation::Status Reset(const std::initializer_list<double>& list);

 public:
  Data() : buf_{nullptr}, size_{0} {}
  Data(Data&&) = default;
  Data& operator=(Data&& other) = default;

  // Not allowed.
  Data(const Data&) = delete;
  Data& operator=(const Data&) = delete;

 private:
  std::unique_ptr<double[]> buf_;
  std::size_t size_;
};

}  // namespace mlvm::array

#endif