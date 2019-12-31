#include "mlvm/Array/Array.h"

#include "gtest/gtest.h"

#include "mlvm/Array/ArrayLike.h"
#include "mlvm/Foundation/StatusOr.h"

namespace mlvm::array {

namespace {

using namespace foundation;

class ArrayTest : public ::testing::Test {};

template <class T>
void inline ASSERT_STATUS_MESSAGE(const StatusOr<T>& status_or,
                                  const std::string& sub_msg) {
  auto err_msg = status_or.statusOrDie().message().value();
  if (err_msg.find(sub_msg) == std::string::npos) {
    FAIL() << "Expected to find: " << sub_msg << "\nBut got: " << err_msg
           << "\n";
  }
}

TEST_F(ArrayTest, CheckArray) {
  auto arr = ArrayLike({1, 2, 3}, {3}).get().consumeValue();
  ASSERT_STREQ("[<3> {1.000, 2.000, 3.000}]", arr->string().c_str());
}

TEST_F(ArrayTest, CheckArrayShape) {
  auto arr = ArrayLike({1, 2, 3}, {3}).get().consumeValue();
  ASSERT_STREQ("<3>", arr->shape().string().c_str());
}

TEST_F(ArrayTest, CheckInvalidData) {
  auto arr_or = ArrayLike({}, {3}).get();
  ASSERT_FALSE(arr_or.ok());
  ASSERT_STATUS_MESSAGE(arr_or, "Data cannot be empty");
}

TEST_F(ArrayTest, CheckInvalidShape) {
  auto arr_or = ArrayLike({3}, {}).get();
  ASSERT_FALSE(arr_or.ok());
  ASSERT_STATUS_MESSAGE(arr_or, "Empty shape");

  arr_or = ArrayLike({3}, {1, 0}).get();
  ASSERT_FALSE(arr_or.ok());
  ASSERT_STATUS_MESSAGE(arr_or, "Non-positive dim");
}

TEST_F(ArrayTest, CheckSizeMismatch) {
  auto arr_or = ArrayLike({1, 2, 3}, {4}).get();
  ASSERT_FALSE(arr_or.ok());
  ASSERT_STATUS_MESSAGE(arr_or, "mismatch");
}

}  // namespace

}  // namespace mlvm::array
