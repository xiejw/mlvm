#include "gtest/gtest.h"

#include "mlvm/lib/Tensor/Tensor.h"

namespace mlvm {
namespace tensor {

class TensorTest : public ::testing::Test {};

TEST_F(TensorTest, CheckConstructor) {
  auto t0 = Tensor::newConstant("t0", {1, 2}, {1.0, 2.0});
}

TEST_F(TensorTest, CheckDebugString) {
  auto t0 = Tensor::newConstant("t0", {1, 2}, {1.0, 2.0});
  ASSERT_STREQ(R"("t0" <1, 2>)", t0.DebugString().c_str());
}

}  // namespace tensor
}  // namespace mlvm
