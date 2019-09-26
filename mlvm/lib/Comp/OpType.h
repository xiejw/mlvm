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

  static OpType Add = OpType(OpAdd);

  Kind kind() const { return kind_; }

 private:
  explicit OpType(Kind kind) : kind_{kind} {};
  Kind kind_;
};

}  // namespace comp
}  // namespace mlvm

#endif
