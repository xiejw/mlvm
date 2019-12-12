#include "mlvm/Local/Data.h"

#include "gtest/gtest.h"

namespace mlvm::local {

namespace {

class DataTest : public ::testing::Test {};

TEST_F(DataTest, CheckEmptyData) {
  local::Data data{};
  ASSERT_FALSE(data.Allocated());
  ASSERT_EQ(0, data.Size());
}

TEST_F(DataTest, CheckArray) {
  local::Data data{};
  data.Reset(new double[3]{1, 2, 3}, 3);
  ASSERT_TRUE(data.Allocated());
  ASSERT_EQ(3, data.Size());
}

TEST_F(DataTest, CheckInitList) {
  local::Data data{};
  data.Reset({1, 2, 3, 4, 5});
  ASSERT_TRUE(data.Allocated());
  ASSERT_EQ(5, data.Size());
}

}  // namespace

}  // namespace mlvm::local
