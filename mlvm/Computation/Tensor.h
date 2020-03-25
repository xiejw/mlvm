#ifndef MLVM_COMPUTATION_H_
#define MLVM_COMPUTATION_H_

#include <string>
#include <string>

namespace mlvm {

class Array {
 public:
  explicit Array(const char* value) : value_{std::string(value)} {};

  std::string debugString() const { return value_; }

 private:
  std::string value_;
};

} // namespace mlvm

#endif
