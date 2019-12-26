#include <iostream>
#include <iterator>

#include "mlvm/Array/Data.h"

using namespace mlvm::array;

int main(int argc, char** argv) {
  Data data{};
  std::cout << data.DebugString() << "\n";
  std::cout << data.Size() << "\n";

  data.Reset(new double[3]{1, 2, 3}, 3);
  std::cout << data.DebugString() << "\n";
  std::cout << data.Size() << "\n";

  data.Reset({1, 2, 3, 4, 5});
  std::cout << data.DebugString() << "\n";
  std::cout << data.Size() << "\n";
}
