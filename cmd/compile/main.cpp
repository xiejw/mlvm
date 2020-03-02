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

// Forward declaration.
class Instruction;

class TResult : public Tensor {
 public:
  TResult(std::string name, Instruction* src, int output_index)
      : name_{std::move(name)}, src_{src}, output_index_{output_index} {};

  std::string DebugString() const override { return name_; }

  Instruction* srcInstructions() const { return src_; }

  int outputIndex() const { return output_index_; }

 private:
  std::string name_;
  Instruction* src_;
  int output_index_;
};

enum class OpType { Add };

class Instruction {
 public:
  Instruction(std::string name, OpType op, std::vector<Tensor*> inputs)
      : name_{std::move(name)}, op_{op}, inputs_{std::move(inputs)} {};

  // TODO: Use outputs.
  void BuildOutputs() {
    auto result = new TResult{absl::StrCat("%o_{", name_, "}"), this, 0};
    outputs_.emplace_back(result);
  };

  std::string DebugString() const {
    std::stringstream ss{};
    switch (op_) {
      case OpType::Add:
        ss << "Add";
        break;
      default:
        ss << "Unknown Op";
    }
    ss << " (" << inputs_.size() << " inputs and " << outputs_.size()
       << " outputs)";
    return ss.str();
  }

 private:
  std::string name_;
  OpType op_;
  std::vector<Tensor*> inputs_;  // unowned.
  std::vector<std::unique_ptr<Tensor>> outputs_ = {};
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
    ins->BuildOutputs();
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
