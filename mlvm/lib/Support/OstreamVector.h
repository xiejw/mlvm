#ifndef SUPPORT_OSTREAMVECTOR
#define SUPPORT_OSTREAMVECTOR

#include <vector>

namespace mlvm {
namespace support {

template <typename T>
void OutputVector(std::ostream &out, const std::vector<T> &vec) {
  int size = vec.size();
  int i = 0;
  for (auto &item : vec) {
    if (++i != size)
      out << item << ", ";
    else
      out << item;
  }
}

}  // namespace support
}  // namespace mlvm

#endif
