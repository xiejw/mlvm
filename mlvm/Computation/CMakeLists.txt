add_library(Computation
  Program.h
  Function.h
  Function.cpp
  Instruction.h
  Instruction.cpp
  OpCode.h
  TensorLike.h
  TensorLike.cpp
  Tuple.h
  Tuple.cpp
)

set(MLVM_LIBS ${MLVM_LIBS} Computation PARENT_SCOPE)

add_library(TestComputation OBJECT
  TensorLikeTest.cpp
)

target_include_directories(TestComputation PUBLIC ${GTEST_INCLUDES})
set(TEST_MLVM_LIBS ${TEST_MLVM_LIBS} TestComputation PARENT_SCOPE)
