#include "mlvm/Foundation/StatusOr.h"

#include "gtest/gtest.h"

namespace mlvm {

namespace {

class FoundationStatusOrTest : public ::testing::Test {};

TEST_F(FoundationStatusOrTest, CheckOK) {
  StatusOr<std::string> status_or{"Hello"};
  ASSERT_TRUE(status_or.ok());
  ASSERT_STREQ("Hello", status_or.valueOrDie().c_str());
}

TEST_F(FoundationStatusOrTest, CheckError) {
  StatusOr<std::string> status_or{ErrorCode::InvalidArguments};
  ASSERT_FALSE(status_or.ok());
  ASSERT_EQ(ErrorCode::InvalidArguments, status_or.statusOrDie().errorCode());
}

// TEST_F(FoundationStatusOrTest, CheckErrorCode) {
//   auto status = Status(ErrorCode::InvalidArguments, "Hello");
//   ASSERT_EQ(ErrorCode::InvalidArguments, status.errorCode());
// }
//
// TEST_F(FoundationStatusOrTest, CheckErrorMessage) {
//   auto status = Status(ErrorCode::InvalidArguments, "Hello");
//   ASSERT_STREQ("Hello", status.message().c_str());
// }

}  // namespace

}  // namespace mlvm
