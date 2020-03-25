#ifndef MLVM_IR_H_
#define MLVM_IR_H_

#include <string>

namespace mlvm {

class Array {
 public:
  explicit Array(const char* value) : value_{std::string(value)} {};

  std::string debugString() const;

 private:
  std::string value_;
};

}  // namespace mlvm

#endif
