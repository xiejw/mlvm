#include "mlvm/Local/Data.h"

#include "gtest/gtest.h"

namespace mlvm::local {

namespace {

class DataTest : public ::testing::Test {};

TEST_F(DataTest, CheckEmptyData) {
  local::Data data{};
  ASSERT_FALSE(data.IsAllocated());
  ASSERT_EQ(0, data.Size());
  ASSERT_STREQ("{}", data.DebugString().c_str());
}

TEST_F(DataTest, CheckArray) {
  local::Data data{};
  data.Reset(new double[3]{1, 2, 3}, 3);
  ASSERT_TRUE(data.IsAllocated());
  ASSERT_EQ(3, data.Size());
  ASSERT_STREQ("{1.000, 2.000, 3.000}", data.DebugString().c_str());
}

TEST_F(DataTest, CheckInitList) {
  local::Data data{};
  data.Reset({1, 2, 3, 4, 5});
  ASSERT_TRUE(data.IsAllocated());
  ASSERT_EQ(5, data.Size());
  ASSERT_STREQ("{1.000, 2.000, 3.000, 4.000, 5.000}",
               data.DebugString().c_str());
}

}  // namespace

}  // namespace mlvm::local
