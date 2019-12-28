#include "mlvm/Foundation/StatusOr.h"

#include "gtest/gtest.h"

namespace mlvm::foundation {

namespace {

class StatusOrTest : public ::testing::Test {};

TEST_F(StatusOrTest, CheckStatus) {
  auto status_or = StatusOr<std::string>{Status::InvalidArguments()};
  ASSERT_FALSE(status_or.Ok());

  auto status = status_or.StatusOrDie();
  ASSERT_EQ(ErrorCode::INVALID_ARGUMENTS, status.Error());
  ASSERT_FALSE(status.Message().has_value());
}

TEST_F(StatusOrTest, CheckValue) {
  auto status_or = StatusOr<std::string>{"hello"};
  ASSERT_TRUE(status_or.Ok());
  ASSERT_STREQ("hello", status_or.ValueOrDie().c_str());
}

TEST_F(StatusOrTest, ConsumeValue) {
  auto status_or = StatusOr<std::string>{"hello"};
  ASSERT_TRUE(status_or.Ok());

  auto value = status_or.ConsumeValue();
  ASSERT_STREQ("hello", value.c_str());
}

}  // namespace

}  // namespace mlvm::foundation
