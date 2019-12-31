#ifndef MLVM_COMPUTATION_TENSORLIKE_
#define MLVM_COMPUTATION_TENSORLIKE_

#include <memory>
#include <sstream>
#include <string>

#include "mlvm/Array/Array.h"

namespace mlvm::computation {

class TensorLike {
 public:
  enum class Type {
    Array,
  };

 public:
  Type type() const { return type_; }
  const std::string& name() const { return name_; }

  std::string string() const {
    assert(type_ == Type::Array);
    std::stringstream ss;
    ss << "`" << name_ << "`: C@";
    ss << array_->string();
    return ss.str();
  }

 private:
  friend class Function;

  TensorLike(std::string name, std::unique_ptr<array::Array> arr)
      : name_{std::move(name)}, type_{Type::Array} {
    array_.swap(arr);
  };

 private:
  std::string name_;
  Type type_;

  std::unique_ptr<array::Array> array_ = {};
};

}  // namespace mlvm::computation

#endif
