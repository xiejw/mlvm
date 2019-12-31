#include <iostream>
#include <iterator>

#include "mlvm/Array/Array.h"
#include "mlvm/Array/ArrayLike.h"
#include "mlvm/Computation/Computation.h"
#include "mlvm/Foundation/Foundation.h"

using namespace mlvm::array;
using namespace mlvm::computation;

int main(int argc, char** argv) {
  auto arr = ArrayLike({1, 2, 3, 4}, {4, 1}).get().consumeValue();
  std::cout << "Array: " << arr->string() << "\n";

  Program p{"test"};
  auto fn = p.makeFunc("main");
  auto t0 = fn->makeTensor({{1, 2, 3, 4, 5}, {5, 1}});
  if (!t0.ok()) {
    std::cout << t0.consumeStatus().message().value();
  }
  std::cout << fn->string() << "\n";
  // auto ins = fn->makeBinaryInst(OpCode::Add, *t0, *t0).consumeValue();
  // auto outputs = fn->makeTupleInst(ins->getOutputs(0)).consumeValue();
  // fn->setOutput(outputs);

  // auto compiledVersion = compile(p);
  // compiledVersion.execute();
}
