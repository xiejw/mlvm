#include <absl/strings/str_cat.h>
#include <iostream>
#include <memory>
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
  Instruction(std::string name, OpType op, std::vector<Tensor*> inputs)
      : name_{std::move(name)}, op_{op}, inputs_{std::move(inputs)} {};

  std::string DebugString() const {
    switch (op_) {
      case OpType::Add:
        return "Add";
    }
    return "Unknown";
  }

 private:
  std::string name_;
  OpType op_;
  std::vector<Tensor*> inputs_;
};

class Function {
 public:
  Function(std::string name) : name_{name} {};

  Tensor* newConst(TConst&& c) {
    std::unique_ptr<TConst> p{new TConst{std::move(c)}};
    consts_.push_back(std::move(p));
    return consts_.back().get();
  };

  template <typename... T>
  Instruction* newInstruction(T&&... args) {
    int count = instructions_.size();
    auto ins =
        new Instruction(absl::StrCat("ins_", count), std::forward<T>(args)...);
    instructions_.push_back(std::unique_ptr<Instruction>{ins});
    return instructions_.back().get();
  }

  std::string DebugString() const {
    std::stringstream ss{};
    ss << "Function: `" << name_ << "`\n";
    ss << "  Consts:\n";
    for (auto& c : consts_) {
      ss << "    - " << c->DebugString() << "\n";
    }

    ss << "\n";
    ss << "  Instructions:\n";
    for (auto& ins : instructions_) {
      ss << "    - " << ins->DebugString() << "\n";
    }

    return ss.str();
  };

 private:
  std::string name_;
  std::vector<std::unique_ptr<TConst>> consts_;
  std::vector<std::unique_ptr<Instruction>> instructions_;
};

}  // namespace mlvm

int main() {
  mlvm::Function fn{"main"};
  auto c0 = fn.newConst(mlvm::TConst{mlvm::Array{"c0"}});

  fn.newInstruction(mlvm::OpType::Add, std::vector{c0, c0});

  std::cout << "const " << c0->DebugString() << "\n";
  std::cout << "func " << fn.DebugString() << "\n";
}
