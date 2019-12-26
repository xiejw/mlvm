#include "mlvm/Array/Data.h"

#include <iomanip>
#include <sstream>

namespace mlvm::array {

std::string Data::DebugString() const {
  std::stringstream ss;
  ss << std::fixed << std::setprecision(3);
  ss << "{";
  for (int i = 0; i < size_; i++) {
    ss << buf_[i];
    if (i != size_ - 1) ss << ", ";
  }
  ss << "}";
  return ss.str();
}

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
