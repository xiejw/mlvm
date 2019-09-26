#ifndef MLVM_COMP_INSTRUCTION
#define MLVM_COMP_INSTRUCTION

#include <memory>
#include <string>
#include <vector>

#include "mlvm/lib/Comp/OpType.h"

namespace mlvm {
namespace comp {

struct Instruction {
  std::string name;
  OpType op_type;
};

}  // namespace comp
}  // namespace mlvm

#endif
