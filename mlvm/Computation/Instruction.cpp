#include "mlvm/Computation/Instruction.h"

namespace mlvm::computation {

Instruction::Instruction(std::string name, OpCode op, std::vector<TensorLike*>&& inputs)
    : name_{std::move(name)}, opCode_{op}, inputs_{inputs}, outputs_{} {}

}  // namespace mlvm::computation
