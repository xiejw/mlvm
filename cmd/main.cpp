#include <iostream>
#include <iterator>

#include "mlvm/Array/Array.h"
#include "mlvm/Array/ArrayLike.h"
#include "mlvm/Computation/Computation.h"
#include "mlvm/Foundation/Foundation.h"

using namespace mlvm::array;
using namespace mlvm::computation;
using namespace mlvm::foundation;

StatusOr<Program> buildProgram() {
  Program p{"test"};
  auto fn = p.makeFunc("main");
  MLVM_ASSIGN_OR_RETURN(t0, fn->makeTensor({{1, 2, 3, 4, 5}, {5, 1}}));
  std::cout << fn->string() << "\n";
  // auto ins = fn->makeBinaryInst(OpCode::Add, *t0, *t0).consumeValue();
  // auto outputs = fn->makeTupleInst(ins->getOutputs(0)).consumeValue();
  // fn->setOutput(outputs);

  // auto compiledVersion = compile(p);
  // compiledVersion.execute();
  return p;
}

int main(int argc, char** argv) {
  auto arr = ArrayLike({1, 2, 3, 4}, {4, 1}).get().consumeValue();
  std::cout << "Array: " << arr->string() << "\n";
  auto p_or = buildProgram();
  if (!p_or.ok()) {
    std::cout << "Failed to build program: "
              << p_or.statusOrDie().message().value();
  }
}
