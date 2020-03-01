#include <iostream>
#include <string>

namespace mlvm {

class Array {
 public:
  explicit Array(const char* value) : value_{std::string(value)} {};

  std::string DebugString() const { return value_; }

 private:
  std::string value_;
};

class Tensor {
  virtual std::string DebugString() const = 0;
};

class TConst : public Tensor {
 public:
  explicit TConst(Array arr) : arr_{std::move(arr)} {};

  std::string DebugString() const override { return arr_.DebugString(); }

 private:
  Array arr_;
};

class TResult : public Tensor {};

enum class OpType { Add };

class Instruction {
 public:
  std::string DebugString() const {
    switch (op_) {
      case OpType::Add:
        return "Add";
    }
    return "Unknown";
  }

 private:
  OpType op_;
};

}  // namespace mlvm

int main() {
  mlvm::TConst c0{mlvm::Array{"c0"}};
  std::cout << "const " << c0.DebugString() << "\n";
}
