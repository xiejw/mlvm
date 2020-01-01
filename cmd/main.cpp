#include <memory>
#include <iostream>
#include <iterator>

#include "mlvm/Array/Array.h"
#include "mlvm/Computation/Computation.h"
#include "mlvm/Foundation/Foundation.h"

using namespace mlvm::array;
using namespace mlvm::computation;
using namespace mlvm::foundation;

StatusOr<Program> buildProgram() {
  Program p{"test"};
  auto fn = p.makeFunc("main");

  MLVM_ASSIGN_OR_RETURN(t0, fn->makeTensor({{1, 2, 3, 4, 5}, {5, 1}}));
  MLVM_CHECK(t0->parentFunc() == fn);

  MLVM_ASSIGN_OR_RETURN(ins, fn->makeBinaryInst(OpCode::Add, t0, t0));
  MLVM_CHECK(1 == ins->outputsCount());

  auto tuple = std::unique_ptr<Tuple>{new Tuple{{ins->getOutput(0)}}};
  MLVM_CHECK(1 == tuple->items().size());
  fn->setOutput(std::move(tuple));

  std::cout << fn->string() << "\n";

  // auto compiledVersion = compile(p);
  // compiledVersion.execute();
  return p;
}

int main(int argc, char** argv) {
  auto p_or = buildProgram();
  if (!p_or.ok()) {
    std::cout << "Failed to build program: "
              << p_or.statusOrDie().message().value();
  }
}
