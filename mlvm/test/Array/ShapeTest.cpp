#include "gtest/gtest.h"

#include "mlvm/lib/Array/Shape.h"

namespace mlvm {
namespace array {

class ShapeTest : public ::testing::Test {};

TEST_F(ShapeTest, CheckDebugString) {
  Shape shape{1, 2};
  ASSERT_STREQ("<1, 2>", shape.DebugString().c_str());
}

TEST_F(ShapeTest, CheckCopy) {
  Shape a{1, 2};
  Shape b{a};
  ASSERT_EQ(&a.dims(), &b.dims());
}

}  // namespace array
}  // namespace mlvm
