#include "gtest/gtest.h"

namespace {

class FoundationStatusOrTest : public ::testing::Test {};

TEST_F(FoundationStatusOrTest, CheckStatus) { ASSERT_TRUE(2 > 1); }

}  // namespace

int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
