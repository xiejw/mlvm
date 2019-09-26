#include <iomanip>
#include <sstream>

#include "mlvm/lib/Support/OstreamVector.h"
#include "mlvm/lib/Tensor/Data.h"

namespace mlvm {
namespace tensor {

std::string Data::DebugString() const {
  std::stringstream ss;
  ss << *this;
  return ss.str();
}

std::ostream &operator<<(std::ostream &out, const Data &s) {
  out << std::fixed << std::setprecision(3) << "[";
  support::OutputVector(out, s.data_);
  out << "]";
  return out;
}

}  // namespace tensor
}  // namespace mlvm
