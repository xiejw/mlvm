add_library(Local
  Data.h
  Data.cpp
)

set(MLVM_LIBS ${MLVM_LIBS} Local PARENT_SCOPE)

add_library(TestLocal OBJECT
  DataTest.cpp
)

target_include_directories(TestLocal PUBLIC ${GTEST_INCLUDES})
set(TEST_MLVM_LIBS ${TEST_MLVM_LIBS} TestLocal PARENT_SCOPE)
