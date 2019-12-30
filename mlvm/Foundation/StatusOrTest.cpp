#include "mlvm/Foundation/StatusOr.h"

#include "gtest/gtest.h"

namespace mlvm::foundation {

namespace {

class StatusOrTest : public ::testing::Test {};

TEST_F(StatusOrTest, CheckStatus) {
  auto status_or = StatusOr<std::string>{Status::InvalidArguments()};
  ASSERT_FALSE(status_or.ok());

  auto status = status_or.StatusOrDie();
  ASSERT_EQ(ErrorCode::InvalidArguments, status.errorCode());
  ASSERT_FALSE(status.message().has_value());
}

TEST_F(StatusOrTest, CheckValue) {
  auto status_or = StatusOr<std::string>{"hello"};
  ASSERT_TRUE(status_or.ok());
  ASSERT_STREQ("hello", status_or.ValueOrDie().c_str());
}

TEST_F(StatusOrTest, ConsumeValue) {
  auto status_or = StatusOr<std::string>{"hello"};
  ASSERT_TRUE(status_or.ok());

  auto value = status_or.ConsumeValue();
  ASSERT_STREQ("hello", value.c_str());
}

}  // namespace

}  // namespace mlvm::foundation
