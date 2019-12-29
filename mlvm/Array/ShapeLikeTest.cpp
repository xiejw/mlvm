#include "mlvm/Array/ShapeLike.h"

#include "gtest/gtest.h"

namespace mlvm::array {

namespace {

class ShapeLikeTest : public ::testing::Test {};

TEST_F(ShapeLikeTest, CheckListConstructor) {
  auto shape_or = ShapeLike({12, 3}).Release();
  ASSERT_TRUE(shape_or.Ok());
}

TEST_F(ShapeLikeTest, CheckMoveConstructor) {
  auto shape_1 = ShapeLike({12, 3}).Release().ConsumeValue();
  auto shape_2 = ShapeLike(std::move(shape_1)).Release().ConsumeValue();
  ASSERT_EQ(36, shape_2.ElementSize());
}

TEST_F(ShapeLikeTest, InvalidEmptyShape) {
  auto shape_or = ShapeLike({}).Release();
  ASSERT_FALSE(shape_or.Ok());
}

TEST_F(ShapeLikeTest, InvalidDim) {
  auto shape_or = ShapeLike({1, 0}).Release();
  ASSERT_FALSE(shape_or.Ok());
}

}  // namespace

}  // namespace mlvm::array
