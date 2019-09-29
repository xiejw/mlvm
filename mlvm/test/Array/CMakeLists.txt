add_library(ArrayTest OBJECT
  DataTest.cpp
  ShapeTest.cpp
  ArrayTest.cpp
)

set(TEST_LIBS ${TEST_LIBS} ArrayTest PARENT_SCOPE)
target_include_directories(ArrayTest PUBLIC ${GTEST_INCLUDES})
