#include <iostream>
#include <memory>
#include <sstream>

#include <string>
#include <vector>

#include <absl/strings/str_cat.h>

#include "mlvm/IR/Function.h"

namespace mlvm {}  // namespace mlvm

int main() {
  mlvm::Function fn{"main"};
  auto c0 = fn.newConst(mlvm::ConstTensor{mlvm::Array{"c0"}});

  auto ins = fn.newInstruction(mlvm::OpType::Add, std::vector{c0, c0});
  auto o0 = ins->outputAt(0);
  ins = fn.newInstruction(mlvm::OpType::Add, std::vector{o0, o0});

  auto status = fn.setOutput(ins->outputAt(0));
  if (!status.ok()) return -1;

  std::cout << "Func:\n" << fn.debugString() << "\n";
}
