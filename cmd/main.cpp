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

  Tuple tuple{};
  tuple.Add({1, 2});

  // std::cout << "Shape: " << shape.ToString() << "\n";
}
