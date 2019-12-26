#include "mlvm/Array/Shape.h"

#include "gtest/gtest.h"

namespace mlvm::array {

namespace {

class ShapeTest : public ::testing::Test {};

TEST_F(ShapeTest, CheckRankAndString) {
  Shape shape{12, 3};
  ASSERT_EQ(2, shape.Rank());
  ASSERT_STREQ("<12, 3>", shape.DebugString().c_str());
}

}  // namespace

}  // namespace mlvm::array
