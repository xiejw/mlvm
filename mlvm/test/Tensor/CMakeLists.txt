add_library(TensorTest OBJECT
  DataTest.cpp
  ShapeTest.cpp
  TensorTest.cpp
)

set(TEST_LIBS ${TEST_LIBS} TensorTest PARENT_SCOPE)
target_include_directories(TensorTest PUBLIC ${GTEST_INCLUDES})
