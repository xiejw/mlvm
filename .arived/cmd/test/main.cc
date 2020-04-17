#include "gtest/gtest.h"

// All test cases should be linked in cmake as OBJECT library.
int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
