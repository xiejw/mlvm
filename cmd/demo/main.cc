#include <absl/strings/str_cat.h>
#include <iostream>
#include <memory>
#include <sstream>
#include <string>
#include <vector>

#include "mlvm/IR/Tensor.h"

namespace mlvm {

class Tensor {
 public:
  virtual ~Tensor(){};
  virtual std::string debugString() const = 0;
};

class ConstTensor : public Tensor {
 public:
  explicit ConstTensor(Array arr) : arr_{std::move(arr)} {};
  ~ConstTensor() override{};

  std::string debugString() const override { return arr_.debugString(); }

 private:
  Array arr_;
};

// Forward declaration.
class Instruction;

class OutputTensor : public Tensor {
 public:
  OutputTensor(std::string name, Instruction* src, int output_index)
      : name_{std::move(name)}, src_{src}, output_index_{output_index} {};

  std::string debugString() const override { return name_; }

  Instruction* srcInstruction() const { return src_; }

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
    auto result = new OutputTensor{absl::StrCat("%o_{", name_, "}"), this, 0};
    outputs_.emplace_back(result);
  };

  std::string_view name() const { return name_; }

  Tensor* outputAt(int i) const { return outputs_[i].get(); }

  std::string debugString() const {
    std::stringstream ss{};
    switch (op_) {
      case OpType::Add:
        ss << "Add";
        break;
      default:
        ss << "Unknown Op";
    }
    ss << " `" << name_ << "`";
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

  Tensor* newConst(ConstTensor&& c) {
    std::unique_ptr<ConstTensor> p{new ConstTensor{std::move(c)}};
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

  // TODO: status
  void setOutput(Tensor* o) { outputs_.push_back(o); }

  std::string debugString() const {
    std::stringstream ss{};
    ss << "Function: `" << name_ << "`\n";
    ss << "  Consts:\n";
    for (auto& c : consts_) {
      ss << "    - " << c->debugString() << "\n";
    }

    ss << "\n";
    ss << "  Instructions:\n";
    for (auto& ins : instructions_) {
      ss << "    - " << ins->debugString() << "\n";
    }

    ss << "\n";
    ss << "  Outputs:\n";
    for (auto& o : outputs_) {
      ss << "    - " << o->debugString() << "\n";
    }

    return ss.str();
  };

 private:
  std::string name_;
  std::vector<std::unique_ptr<ConstTensor>> consts_ = {};
  std::vector<std::unique_ptr<Instruction>> instructions_ = {};
  std::vector<Tensor*> outputs_ = {};
};

}  // namespace mlvm

int main() {
  mlvm::Function fn{"main"};
  auto c0 = fn.newConst(mlvm::ConstTensor{mlvm::Array{"c0"}});

  auto ins = fn.newInstruction(mlvm::OpType::Add, std::vector{c0, c0});
  auto o0 = ins->outputAt(0);
  ins = fn.newInstruction(mlvm::OpType::Add, std::vector{o0, o0});

  fn.setOutput(ins->outputAt(0));

  std::cout << "Func:\n" << fn.debugString() << "\n";
}
