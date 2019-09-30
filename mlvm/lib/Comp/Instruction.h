#ifndef MLVM_COMP_INSTRUCTION
#define MLVM_COMP_INSTRUCTION

#include <string>
#include <vector>

#include "mlvm/lib/Comp/OpType.h"
#include "mlvm/lib/Comp/Tensor.h"

namespace mlvm {
namespace comp {

struct Instruction {
  std::string name;
  OpType op_type;
  std::vector<Tensor*> operands;

 public:
  friend std::ostream& operator<<(std::ostream& os, const Instruction& ins);
  std::string DebugString() const;
};

}  // namespace comp
}  // namespace mlvm

#endif
