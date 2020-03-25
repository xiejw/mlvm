#include <iostream>
#include <memory>
#include <sstream>

#include <string>
#include <vector>

#include <absl/strings/str_cat.h>

#include "mlvm/IR/Instruction.h"
#include "mlvm/IR/Tensor.h"

namespace mlvm {

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
