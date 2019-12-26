#include "mlvm/Array/Data.h"

#include "gtest/gtest.h"

namespace mlvm::array {

namespace {

class DataTest : public ::testing::Test {};

TEST_F(DataTest, CheckEmptyData) {
  Data data{};
  ASSERT_FALSE(data.IsAllocated());
  ASSERT_EQ(0, data.Size());
  ASSERT_STREQ("{}", data.ToString().c_str());
}

TEST_F(DataTest, CheckArray) {
  Data data{};
  data.Reset(new double[3]{1, 2, 3}, 3);
  ASSERT_TRUE(data.IsAllocated());
  ASSERT_EQ(3, data.Size());
  ASSERT_STREQ("{1.000, 2.000, 3.000}", data.ToString().c_str());
}

TEST_F(DataTest, CheckInitList) {
  Data data{};
  data.Reset({1, 2, 3, 4, 5});
  ASSERT_TRUE(data.IsAllocated());
  ASSERT_EQ(5, data.Size());
  ASSERT_STREQ("{1.000, 2.000, 3.000, 4.000, 5.000}", data.ToString().c_str());
}

}  // namespace

}  // namespace mlvm::array
