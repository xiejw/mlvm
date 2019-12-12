#include "mlvm/Local/Data.h"

namespace mlvm::local {

std::string Data::DebugString() const { return "Hello from Data"; }

void Data::Reset(double* new_data, std::size_t size) {
  size_ = size;
  buf_.reset(new_data);
}

void Data::Reset(const std::initializer_list<double>& list) {
  auto size = list.size();
  auto new_data = new double[size];

  int i = 0;
  for (auto& el : list) {
    new_data[i++] = el;
  }
  Reset(new_data, size);
}

}  // namespace mlvm::local
