#include "mlvm/Computation/Instruction.h"

#include "gtest/gtest.h"

#include "mlvm/Array/ArrayLike.h"
#include "mlvm/Computation/Function.h"
#include "mlvm/Computation/TensorLike.h"
#include "mlvm/Foundation/StatusOr.h"

namespace mlvm::computation {

namespace {

class InstructionTest : public ::testing::Test {};

TEST_F(InstructionTest, CheckMemeberFuncs) {
  Function fn{"test"};
  TensorLike* c = fn.makeTensor({{1, 2, 3, 4, 5}, {5, 1}}).consumeValue();

  Instruction ins{"i", OpCode::Add, {c, c}, &fn};
  ASSERT_STREQ("i", ins.name().c_str());
  ASSERT_EQ(OpCode::Add, ins.opCode());
  ASSERT_EQ(1, ins.outputsCount());
}

}  // namespace
}  // namespace mlvm::computation
