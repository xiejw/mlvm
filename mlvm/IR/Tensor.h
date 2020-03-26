#ifndef MLVM_IR_TENSOR_H_
#define MLVM_IR_TENSOR_H_

#include <string>

namespace mlvm {

class Array {
 public:
  explicit Array(const char* value) : value_{std::string(value)} {};

  std::string debugString() const;

 private:
  std::string value_;
};

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

  std::string debugString() const override;

  Instruction* srcInstruction() const { return src_; }

  int outputIndex() const { return output_index_; }

 private:
  std::string name_;
  Instruction* src_;
  int output_index_;
};

}  // namespace mlvm

#endif
