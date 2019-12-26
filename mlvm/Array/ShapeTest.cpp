#include "mlvm/Array/Shape.h"

#include "gtest/gtest.h"

namespace mlvm::array {

namespace {

class ShapeTest : public ::testing::Test {};

TEST_F(ShapeTest, CheckRankAndString) {
  Shape shape{12, 3};
  ASSERT_EQ(2, shape.Rank());
  ASSERT_STREQ("<12, 3>", shape.ToString().c_str());
}

TEST_F(ShapeTest, InvalidEmptyShape) {
  try {
    Shape shape{};
    FAIL() << "Should not reach here.";
  } catch (const char* msg) {
    ASSERT_STREQ("Empty shape is not allowed.", msg);
  }
}

TEST_F(ShapeTest, InvalidDim) {
  try {
    Shape shape{1, 0};
    FAIL() << "Should not reach here.";
  } catch (const char* msg) {
    ASSERT_STREQ("Non-positive dim is not allowed.", msg);
  }
}

}  // namespace

}  // namespace mlvm::array
