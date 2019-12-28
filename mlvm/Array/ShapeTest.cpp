#include "mlvm/Array/Shape.h"

#include "gtest/gtest.h"

namespace mlvm::array {

namespace {

class ShapeTest : public ::testing::Test {};

TEST_F(ShapeTest, CheckRankAndString) {
  auto shape_or = Shape::New({12, 3});
  ASSERT_TRUE(shape_or.Ok());

  auto shape = shape_or.ValueOrDie();
  ASSERT_EQ(2, shape.Rank());
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
