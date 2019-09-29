#include "gtest/gtest.h"

#include "mlvm/lib/Array/Array.h"

namespace mlvm {
namespace tensor {

class ArrayTest : public ::testing::Test {};

TEST_F(ArrayTest, CheckConstructor) {
  auto t0 = Array("t0", {1, 2}, {1.0, 2.0});
}

TEST_F(ArrayTest, CheckDebugString) {
  auto t0 = Array("t0", {1, 2}, {1.0, 2.0});
  ASSERT_STREQ(R"("t0" <1, 2>)", t0.DebugString().c_str());
}

TEST_F(ArrayTest, CheckData) {
  auto t0 = Array("t0", {1, 2}, {1.0, 2.0});
  ASSERT_STREQ("[1.000, 2.000]", t0.data().DebugString().c_str());
}

}  // namespace tensor
}  // namespace mlvm
