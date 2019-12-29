#include "mlvm/Array/Shape.h"
#include "mlvm/Array/ShapeLike.h"

#include "gtest/gtest.h"

namespace mlvm::array {

namespace {

class ShapeTest : public ::testing::Test {};

TEST_F(ShapeTest, CheckElementSize) {
  auto shape = ShapeLike({12, 3}).Release().ConsumeValue();
  ASSERT_EQ(36, shape.ElementSize());
}

TEST_F(ShapeTest, CheckRank) {
  auto shape = ShapeLike({12, 3}).Release().ConsumeValue();
  ASSERT_EQ(2, shape.Rank());
}

TEST_F(ShapeTest, CheckString) {
  auto shape = ShapeLike({12, 3}).Release().ConsumeValue();
  ASSERT_STREQ("<12, 3>", shape.ToString().c_str());
}

}  // namespace

}  // namespace mlvm::array
