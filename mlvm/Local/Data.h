#ifndef MLVM_LOCAL_DATA_
#define MLVM_LOCAL_DATA_

#include <cstddef>
#include <initializer_list>
#include <memory>
#include <string>

namespace mlvm::local {

// Represents a data buffer.
class Data {
 public:
  std::string DebugString() const;

  bool Allocated() const { return size_ > 0; }
  int Size() const { return size_; }

  // Move the `new_data` into this strucuture. So, owns the `new_data`.
  void Reset(double* new_data, std::size_t size);

  // Copy the `list` into this strucuture.
  void Reset(std::initializer_list<double> list);

 private:
  std::unique_ptr<double[]> buf_;
  int size_;
};

}  // namespace mlvm::local

#endif
