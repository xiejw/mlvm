#include <iostream>
#include <iterator>

#include "mlvm/Array/Array.h"
#include "mlvm/Computation/Computation.h"
#include "mlvm/Foundation/Foundation.h"

using namespace mlvm::array;
using namespace mlvm::computation;

int main(int argc, char** argv) {
  auto arr = Array::New({1, 2, 3, 4, 5}, {4, 1}).consumeValue();
  std::cout << "Array: " << arr->string() << "\n";

  // Program p{"test"};
  // auto fn = p.makeFunc("main");
  // auto t0 = fn->makeTensor({1, 2, 3, 4, 5}, {4, 1}).consumeValue();
  // auto ins = fn->addBinaryInst(OpCode::Add, *t0, *t0).consumeValue();
  // auto outputs = fn->addTupleInst(ins->getOutputs(0)).consumeValue();
  // fn->setOutput(outputs);

  // auto compiledVersion = compile(p);
  // compiledVersion.execute();
}
