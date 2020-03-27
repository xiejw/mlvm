#ifndef MLVM_IR_FUNCTION_H_
#define MLVM_IR_FUNCTION_H_

#include <absl/strings/str_cat.h>

#include "mlvm/Foundation/Status.h"
#include "mlvm/IR/Instruction.h"
#include "mlvm/IR/Tensor.h"

namespace mlvm::IR {

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
    MLVM_FATAL_IF_ERROR(ins->buildOutputs());
    instructions_.push_back(std::unique_ptr<Instruction>{ins});
    return instructions_.back().get();
  }

  Status setOutput(Tensor* o) {
    outputs_.push_back(o);
    return Status::OK;
  }

  std::string debugString() const;
  std::string_view name() const { return name_; }

 private:
  std::string name_;
  std::vector<std::unique_ptr<ConstTensor>> consts_ = {};
  std::vector<std::unique_ptr<Instruction>> instructions_ = {};
  std::vector<Tensor*> outputs_ = {};
};

}  // namespace mlvm::IR

#endif
