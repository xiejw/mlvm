#include "mlvm/Foundation/StatusOr.h"

#include "gtest/gtest.h"

namespace mlvm {

namespace {

class FoundationStatusOrTest : public ::testing::Test {};

TEST_F(FoundationStatusOrTest, CheckOK) {
  StatusOr<std::string> status{"Hello"};
  ASSERT_TRUE(status.ok());
}

TEST_F(FoundationStatusOrTest, CheckError) {
  StatusOr<std::string> status{ErrorCode::InvalidArguments};
  ASSERT_FALSE(status.ok());
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
