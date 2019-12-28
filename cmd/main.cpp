#include <iostream>
#include <iterator>

#include "mlvm/Array/Array.h"
#include "mlvm/Foundation/Status.h"

using namespace mlvm::array;

int main(int argc, char** argv) {
  auto arr = Array::New({1, 2, 3, 4, 5}, {4, 1}).ConsumeValue();
  std::cout << "Array: " << arr.ToString() << "\n";

  // Shape shape{1, 2};
  // std::cout << "Shape: " << shape.ToString() << "\n";
}
