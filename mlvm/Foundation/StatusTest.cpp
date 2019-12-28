#include "mlvm/Foundation/Status.h"

#include "gtest/gtest.h"

namespace mlvm::foundation {

namespace {

class StatusTest : public ::testing::Test {};

TEST_F(StatusTest, CheckOK) {
  auto status = Status::OK;
  ASSERT_TRUE(status.Ok());
}

TEST_F(StatusTest, CheckErrorCode) {
  auto status = Status(ErrorCode::INVALID_ARGUMENTS, "Hello");
  ASSERT_EQ(ErrorCode::INVALID_ARGUMENTS, status.Error());
}

TEST_F(StatusTest, CheckErrorMessage) {
  auto status = Status(ErrorCode::INVALID_ARGUMENTS, "Hello");
  ASSERT_STREQ("Hello", status.Message().value().c_str());

  status = Status(ErrorCode::INVALID_ARGUMENTS);
  ASSERT_FALSE(status.Message().has_value());
}

}  // namespace

}  // namespace mlvm::foundation
