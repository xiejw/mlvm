#include "mlvm/Computation/Function.h"
#include "mlvm/Computation/TensorLike.h"

#include "gtest/gtest.h"

#include "mlvm/Array/ArrayLike.h"
#include "mlvm/Foundation/StatusOr.h"

namespace mlvm::computation {

namespace {

using namespace foundation;

class TensorLikeTest : public ::testing::Test {};

TEST_F(TensorLikeTest, CheckArrayTensor) {
  Function fn{"test"};
  auto tensor_or = fn.makeTensor({{1, 2, 3, 4, 5}, {5, 1}});
  ASSERT_TRUE(tensor_or.ok());
}

}  // namespace
}  // namespace mlvm::computation
