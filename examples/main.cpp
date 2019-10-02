#include <iostream>
#include <iomanip>

#include "gflags/gflags.h"

#include "mlvm/lib/Comp/Computation.h"
#include "mlvm/lib/Array/Array.h"

using mlvm::array::Array;
using mlvm::array::Shape;
using mlvm::comp::OpType;

int main(int argc, char* argv[]) {
  gflags::ParseCommandLineFlags(&argc, &argv, true);

  auto t1 = Array("t1", {1,2}, {2.12, 3.13});
  std::cout << "Hello " << t1.data() << "\n";

  auto comp = mlvm::comp::Computation();
  comp.newInstruction("ins", OpType::Add(), {});
  // comp.newInstruction("ins", OpType::Add(), {t1, t1});
  std::cout << "Comp: " << comp.DebugString();
}
