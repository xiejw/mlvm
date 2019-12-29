#include <iostream>
#include <iterator>

#include "mlvm/Array/Array.h"
#include "mlvm/Computation/Tuple.h"
#include "mlvm/Foundation/Status.h"

using namespace mlvm::array;
using namespace mlvm::computation;

int main(int argc, char** argv) {
  auto arr = Array::New({1, 2, 3, 4, 5}, {4, 1}).ConsumeValue();
  std::cout << "Array: " << arr.ToString() << "\n";

  Program p{"test"};
  auto fn = p.makeFunc("main");
  TensorLike lhs{}, rhs{};
  auto ins = fn->addBinaryInst(OpCode::Add, lhs, rhs).ConsumeValue();
  auto outputs = fn->addTupleInst(ins->getOutputs(0)).ConsumeValue();
  fn->setOutput(outputs);

  // auto compiledVersion = compile(p);
  // compiledVersion.execute();
}
