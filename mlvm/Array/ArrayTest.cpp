#include "mlvm/Array/Array.h"

#include "gtest/gtest.h"

namespace mlvm::array {

namespace {

using namespace foundation;

class ArrayTest : public ::testing::Test {};

void inline ASSERT_STATUS_MESSAGE(const StatusOr<Array>& status_or,
                                  const std::string& sub_msg) {
  auto err_msg = status_or.StatusOrDie().message().value();
  if (err_msg.find(sub_msg) == std::string::npos) {
    FAIL() << "Expected to find: " << sub_msg << "\nBut got: " << err_msg
           << "\n";
  }
}

TEST_F(ArrayTest, CheckArray) {
  auto arr = Array::New({1, 2, 3}, {3}).ConsumeValue();
  ASSERT_STREQ("[<3> {1.000, 2.000, 3.000}]", arr.string().c_str());
}

TEST_F(ArrayTest, CheckInvalidData) {
  auto arr_or = Array::New({}, {3});
  ASSERT_FALSE(arr_or.ok());
  ASSERT_STATUS_MESSAGE(arr_or, "Data cannot be empty");
}

TEST_F(ArrayTest, CheckInvalidShape) {
  auto arr_or = Array::New({3}, {});
  ASSERT_FALSE(arr_or.ok());
  ASSERT_STATUS_MESSAGE(arr_or, "Empty shape");

  arr_or = Array::New({3}, {1, 0});
  ASSERT_FALSE(arr_or.ok());
  ASSERT_STATUS_MESSAGE(arr_or, "Non-positive dim");
}

TEST_F(ArrayTest, CheckSizeMismatch) {
  auto arr_or = Array::New({1, 2, 3}, {4});
  ASSERT_FALSE(arr_or.ok());
  ASSERT_STATUS_MESSAGE(arr_or, "mismatch");
}

}  // namespace

}  // namespace mlvm::array
