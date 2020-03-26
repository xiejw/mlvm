#include "mlvm/Foundation/Status.h"

#include "gtest/gtest.h"

namespace mlvm {

namespace {

class FoundationStatusTest : public ::testing::Test {};

TEST_F(FoundationStatusTest, CheckOK) {
  auto status = Status::OK;
  ASSERT_TRUE(status.ok());
}

TEST_F(FoundationStatusTest, CheckErrorCode) {
  auto status = Status(ErrorCode::InvalidArguments, "Hello");
  ASSERT_EQ(ErrorCode::InvalidArguments, status.errorCode());
}

TEST_F(FoundationStatusTest, CheckErrorMessage) {
  auto status = Status(ErrorCode::InvalidArguments, "Hello");
  ASSERT_STREQ("Hello", status.message().c_str());
}

}  // namespace

}  // namespace mlvm
