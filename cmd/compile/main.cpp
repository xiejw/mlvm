#include <iostream>
#include <sstream>
#include <string>
#include <vector>

namespace mlvm {

class Array {
 public:
  explicit Array(const char* value) : value_{std::string(value)} {};

  std::string DebugString() const { return value_; }

 private:
  std::string value_;
};

class Tensor {
 public:
  virtual ~Tensor(){};
  virtual std::string DebugString() const = 0;
};

class TConst : public Tensor {
 public:
  explicit TConst(Array arr) : arr_{std::move(arr)} {};
  ~TConst() override{};

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

class Function {
 public:
  Function(std::string name) : name_{name} {};

  Tensor* newConst(TConst&& c) {
    consts_.push_back(std::move(c));
    return &consts_.back();
  };

  std::string DebugString() const {
    std::stringstream ss{};
    ss << "Function: `" << name_ << "`\n";
    ss << "  Consts:\n";
    for (auto& c : consts_) {
      ss << "    - " << c.DebugString() << "\n";
    }
    return ss.str();
  };

 private:
  std::string name_;
  std::vector<TConst> consts_;
};

}  // namespace mlvm

int main() {
  mlvm::Function fn{"main"};
  auto c0 = fn.newConst(mlvm::TConst{mlvm::Array{"c0"}});
  std::cout << "const " << c0->DebugString() << "\n";
  std::cout << "func " << fn.DebugString() << "\n";
}
