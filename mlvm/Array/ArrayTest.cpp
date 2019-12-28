#include "mlvm/Array/Array.h"

#include "gtest/gtest.h"

namespace mlvm::array {

namespace {

class ArrayTest : public ::testing::Test {};

// TEST_F(ArrayTest, CheckEmptyData) {
//   Data data{};
//   ASSERT_FALSE(data.IsAllocated());
//   ASSERT_EQ(0, data.Size());
//   ASSERT_STREQ("{}", data.ToString().c_str());
// }

TEST_F(ArrayTest, CheckArray) {
  auto arr = Array::New({1, 2, 3}, {3}).ConsumeValue();
  ASSERT_STREQ("[<3> {1.000, 2.000, 3.000}]", arr.ToString().c_str());
}

// TEST_F(ArrayTest, CheckInitList) {
//   Data data{};
//   data.Reset({1, 2, 3, 4, 5});
//   ASSERT_TRUE(data.IsAllocated());
//   ASSERT_EQ(5, data.Size());
//   ASSERT_STREQ("{1.000, 2.000, 3.000, 4.000, 5.000}", data.ToString().c_str());
// }

}  // namespace

}  // namespace mlvm::array
