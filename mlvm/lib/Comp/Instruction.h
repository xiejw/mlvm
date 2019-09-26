#ifndef MLVM_COMP_INSTRUCTION
#define MLVM_COMP_INSTRUCTION

#include <string>

#include "mlvm/lib/Comp/OpType.h"

namespace mlvm {
namespace comp {

struct Instruction {
  std::string name;
  OpType op_type;

 public:
  std::string DebugString() const;
};

}  // namespace comp
}  // namespace mlvm

#endif
