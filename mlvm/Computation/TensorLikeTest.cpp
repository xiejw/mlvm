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

  auto tensor = tensor_or.consumeValue();
  ASSERT_STREQ("%c_0", tensor->name().c_str());
  ASSERT_EQ(TensorLike::Type::Constant, tensor->type());
  ASSERT_STREQ("<5, 1>", tensor->shape().string().c_str());
  ASSERT_EQ(&fn, tensor->parentFunc());
  ASSERT_EQ(nullptr, tensor->parentIns());
  ASSERT_STREQ("`%c_0`: C@[<5, 1> {1.000, 2.000, 3.000, 4.000, 5.000}]",
               tensor->string().c_str());
}

}  // namespace
}  // namespace mlvm::computation
