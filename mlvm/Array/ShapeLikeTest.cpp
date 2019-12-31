#include "mlvm/Array/ShapeLike.h"

#include "gtest/gtest.h"

namespace mlvm::array {

namespace {

class ShapeLikeTest : public ::testing::Test {};

TEST_F(ShapeLikeTest, CheckListConstructor) {
  auto shape_or = ShapeLike({12, 3}).get();
  ASSERT_TRUE(shape_or.ok());
}

TEST_F(ShapeLikeTest, CheckMoveConstructor) {
  auto shape_1 = ShapeLike({12, 3}).get().consumeValue();
  auto shape_2 = ShapeLike(std::move(shape_1)).get().consumeValue();
  ASSERT_EQ(36, shape_2.elementSize());
}

TEST_F(ShapeLikeTest, CheckCopyConstructor) {
  Shape shape_1 = ShapeLike({12, 3}).get().consumeValue();
  ASSERT_EQ(36, shape_1.elementSize());

  Shape shape_2 = ShapeLike(shape_1).get().consumeValue();
  ASSERT_EQ(36, shape_2.elementSize());
}

TEST_F(ShapeLikeTest, InvalidEmptyShape) {
  auto shape_or = ShapeLike({}).get();
  ASSERT_FALSE(shape_or.ok());
}

TEST_F(ShapeLikeTest, InvalidDim) {
  auto shape_or = ShapeLike({1, 0}).get();
  ASSERT_FALSE(shape_or.ok());
}

}  // namespace

}  // namespace mlvm::array
