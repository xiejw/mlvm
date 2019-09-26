#include <iostream>
#include <iomanip>

#include "gflags/gflags.h"

#include "mlvm/lib/Comp/Computation.h"
#include "mlvm/lib/Tensor/Shape.h"
#include "mlvm/lib/Tensor/Tensor.h"

using mlvm::comp::OpType;

int main(int argc, char* argv[]) {
  gflags::ParseCommandLineFlags(&argc, &argv, true);

  mlvm::tensor::Shape shape {1,2};
  auto t1 = mlvm::tensor::Tensor::newConstant("t1", {1,2}, {2.12, 3.13});
  std::cout << "Hello " << shape << "\n";
  std::cout << "Hello " << t1.data() << "\n";

  auto comp = mlvm::comp::Computation();
  comp.newInstruction("ins", OpType::Add());
  std::cout << "Comp: " << comp.DebugString();
}
