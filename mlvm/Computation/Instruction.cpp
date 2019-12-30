#include "mlvm/Computation/Instruction.h"

namespace mlvm::computation {

Instruction::Instruction(OpCode op, std::vector<TensorLike*>&& inputs)
    : opCode_{op}, inputs_{inputs}, outputs_{} {}

}  // namespace mlvm::computation
