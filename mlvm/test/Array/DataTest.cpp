#include "gtest/gtest.h"

#include "mlvm/lib/Array/Data.h"

namespace mlvm {
namespace array {

class DataTest : public ::testing::Test {};

TEST_F(DataTest, CheckDebugString) {
  Data data{1, 2};
  ASSERT_STREQ("[1.000, 2.000]", data.DebugString().c_str());
}

}  // namespace array
}  // namespace mlvm
