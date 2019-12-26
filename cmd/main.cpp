#include <iostream>
#include <iterator>

#include "mlvm/Array/Array.h"

using namespace mlvm::array;

int main(int argc, char** argv) {
  Data data{};
  data.Reset({1, 2, 3, 4, 5});
  std::cout << "Data: " << data.ToString() << "\n";

  Shape shape{1, 2};
  std::cout << "Shape: " << shape.ToString() << "\n";
}
