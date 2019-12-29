#include "mlvm/Array/Data.h"

#include <iomanip>
#include <sstream>

namespace mlvm::array {

using namespace foundation;

std::string Data::string() const {
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

Status Data::reset(double* new_data, std::size_t size) {
  if (size == 0) return Status::InvalidArguments("Data cannot be empty.");
  if (new_data == nullptr)
    return Status::InvalidArguments("Cannot take nullptr buffer.");

  size_ = size;
  buf_.reset(new_data);
  return Status::OK;
}

Status Data::reset(const std::initializer_list<double>& list) {
  auto size = list.size();
  auto new_data = new double[size];

  int i = 0;
  for (auto& el : list) {
    new_data[i++] = el;
  }
  return reset(new_data, size);
}

}  // namespace mlvm::array
