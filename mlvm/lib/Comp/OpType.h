#ifndef MLVM_COMP_OPTYPE
#define MLVM_COMP_OPTYPE

namespace mlvm {
namespace comp {

class OpType {
 public:
  enum Kind {
    OpAdd,
    OpMul,
  };

 public:
  /// All singletone for OpTypes.
  static OpType& Add() {
    static OpType* singleton = new OpType(OpAdd);
    return *singleton;
  }

 public:
  Kind kind() const { return kind_; }

 private:
  friend std::ostream& operator<<(std::ostream& os, const OpType& op);

 private:
  explicit OpType(Kind kind) : kind_{kind} {};
  Kind kind_;
};

}  // namespace comp
}  // namespace mlvm

#endif
