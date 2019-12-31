add_library(Computation
  Program.h
  Function.h
  Function.cpp
  Instruction.h
  Instruction.cpp
  OpCode.h
  TensorLike.h
  Tuple.h
  Tuple.cpp
)

set(MLVM_LIBS ${MLVM_LIBS} Computation PARENT_SCOPE)

# add_library(TestFoundation OBJECT
#   StatusTest.cpp
#   StatusOrTest.cpp
# )
#
# target_include_directories(TestFoundation PUBLIC ${GTEST_INCLUDES})
# set(TEST_MLVM_LIBS ${TEST_MLVM_LIBS} TestFoundation PARENT_SCOPE)
