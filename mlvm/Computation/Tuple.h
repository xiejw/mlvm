#ifndef MLVM_COMPUTATION_TUPLE_
#define MLVM_COMPUTATION_TUPLE_

#include <memory>
#include <optional>
#include <string>
#include <unordered_map>
#include <vector>

#include "mlvm/Array/ShapeLike.h"
#include "mlvm/Foundation/Status.h"
#include "mlvm/Foundation/StatusOr.h"

namespace mlvm::computation {

using namespace foundation;

struct Item {
  array::Shape shape;
  std::optional<std::string> name;
};

class Tuple {
 public:
  foundation::Status Add(array::ShapeLike shape_like);

 private:
  std::vector<std::unique_ptr<Item>> items_;
};

class TensorLike {};

enum class OpCode {
  Add,
};

class Instruction {
 public:
  Instruction(OpCode op) : opCode_{op} {};

 public:
  const TensorLike& getOutputs(int i) { return outputs_[i]; };

 private:
  OpCode opCode_;
  std::vector<TensorLike*> inputs_;
  std::vector<TensorLike> outputs_;
};

class Function {
 public:
  Function(std::string name) : ins_vec_{}, name_{name} {};

 public:
  StatusOr<Instruction*> addBinaryInst(OpCode op, const TensorLike& lhs,
                                       const TensorLike& rhs) {
    return nullptr;
  };

  StatusOr<Instruction*> addTupleInst(const TensorLike& lhs) {
    return nullptr;
  };

  StatusOr<Instruction*> setOutput(Instruction* const ins) { return nullptr; }

 private:
  std::vector<std::unique_ptr<Instruction>> ins_vec_;
  ;
  std::string name_;
};

class Program {
 public:
  Program(std::string name) : name_{name} {};

 public:
  Function* makeFunc(std::string fn_name) {
    auto fn = new Function{fn_name};
    fns_[fn_name] = std::unique_ptr<Function>{fn};
    // Check key dost noe exist.
    return fns_[fn_name].get();
  }

 private:
  std::unordered_map<std::string, std::unique_ptr<Function>> fns_;
  std::string name_;
};

}  // namespace mlvm::computation

#endif
