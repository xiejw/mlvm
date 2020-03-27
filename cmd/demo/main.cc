#include <iostream>

#include "mlvm/IR/Function.h"
#include "mlvm/Runtime/Evaluator.h"

int main() {
  mlvm::IR::Function fn{"main"};
  auto c0 = fn.newConst(mlvm::IR::ConstTensor{mlvm::IR::Array{"c0"}});
  auto ins = fn.newInstruction(mlvm::IR::OpType::Add, std::vector{c0, c0});
  auto o0 = ins->outputAt(0);
  ins = fn.newInstruction(mlvm::IR::OpType::Add, std::vector{o0, o0});
  MLVM_FATAL_IF_ERROR(fn.setOutput(ins->outputAt(0)));

  std::cout << "Func:\n" << fn.debugString() << "\n";

  mlvm::RT::Evaluator eval{};
  MLVM_FATAL_IF_ERROR(eval.run(fn));
}
