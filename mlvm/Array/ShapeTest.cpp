#include "mlvm/Array/Shape.h"
#include "mlvm/Array/ShapeLike.h"

#include "gtest/gtest.h"

namespace mlvm::array {

namespace {

class ShapeTest : public ::testing::Test {};

TEST_F(ShapeTest, CheckConstructor) {
  auto shape_or = ShapeLike({12, 3}).Release();
  ASSERT_TRUE(shape_or.Ok());
}

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

TEST_F(ShapeTest, InvalidEmptyShape) {
  auto shape_or = ShapeLike({}).Release();
  ASSERT_FALSE(shape_or.Ok());
}

TEST_F(ShapeTest, InvalidDim) {
  auto shape_or = ShapeLike({1, 0}).Release();
  ASSERT_FALSE(shape_or.Ok());
}

}  // namespace

}  // namespace mlvm::array
