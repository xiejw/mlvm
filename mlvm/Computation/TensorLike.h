#ifndef MLVM_COMPUTATION_TENSORLIKE_
#define MLVM_COMPUTATION_TENSORLIKE_

#include <optional>
#include <memory>
#include <string>

#include "mlvm/Array/Array.h"

namespace mlvm::computation {

class Function;
class Instruction;

class TensorLike {
 public:
  enum class Type {
    Constant,
    Output,
  };

 public:
  Type type() const { return type_; }

  const std::string& name() const { return name_; }

  const array::Shape& shape() const;

  Function* parentFunc() const { return parent_fn_; }

  Instruction* parentIns() const { return parent_ins_; }

  std::string string() const;

 private:
  friend class Function;

  // Constant has no parent Instruction
  TensorLike(std::string name, std::unique_ptr<array::Array> arr,
             Function* parent_fn)
      : name_{std::move(name)},
        type_{Type::Constant},
        parent_fn_{parent_fn},
        parent_ins_{nullptr} {
    array_.swap(arr);
  };

  // Sets the output.
  TensorLike(std::string name, array::ShapeLike output_shape, Function* parent_fn, Instruction* parent_ins)
      : name_{std::move(name)},
        type_{Type::Output},
        parent_fn_{parent_fn},
        parent_ins_{parent_ins} {

    output_shape_ = output_shape.get().consumeValue();
  };

 private:
  std::string name_;
  Type type_;

  Function* parent_fn_;
  Instruction* parent_ins_;

  std::optional<array::Shape> output_shape_ = {};
  std::unique_ptr<array::Array> array_ = {};
};

}  // namespace mlvm::computation

#endif
