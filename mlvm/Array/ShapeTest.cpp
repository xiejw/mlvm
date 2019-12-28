#include "mlvm/Array/Shape.h"

#include "gtest/gtest.h"

namespace mlvm::array {

namespace {

class ShapeTest : public ::testing::Test {};

TEST_F(ShapeTest, CheckConstructor) {
  auto shape_or = Shape::New({12, 3});
  ASSERT_TRUE(shape_or.Ok());
}

TEST_F(ShapeTest, CheckElementSize) {
  auto shape = Shape::New({12, 3}).ConsumeValue();
  ASSERT_EQ(36, shape.ElementSize());
}

TEST_F(ShapeTest, CheckRank) {
  auto shape = Shape::New({12, 3}).ConsumeValue();
  ASSERT_EQ(2, shape.Rank());
}

TEST_F(ShapeTest, CheckString) {
  auto shape = Shape::New({12, 3}).ConsumeValue();
  ASSERT_STREQ("<12, 3>", shape.ToString().c_str());
}

TEST_F(ShapeTest, InvalidEmptyShape) {
  auto shape_or = Shape::New({});
  ASSERT_FALSE(shape_or.Ok());
}

TEST_F(ShapeTest, InvalidDim) {
  auto shape_or = Shape::New({1, 0});
  ASSERT_FALSE(shape_or.Ok());
}

}  // namespace

}  // namespace mlvm::array
