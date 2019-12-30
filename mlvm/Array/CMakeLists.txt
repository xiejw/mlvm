add_library(Array
  Data.h
  Data.cpp
  Shape.h
  Shape.cpp
  ShapeLike.h
  ShapeLike.cpp
  Array.h
  ArrayImpl.h
  ArrayImpl.cpp
  ArrayLike.h
  ArrayLike.cpp
)

set(MLVM_LIBS ${MLVM_LIBS} Array PARENT_SCOPE)

add_library(TestArray OBJECT
  DataTest.cpp
  ShapeTest.cpp
  ShapeLikeTest.cpp
  ArrayImplTest.cpp
)

target_include_directories(TestArray PUBLIC ${GTEST_INCLUDES})
set(TEST_MLVM_LIBS ${TEST_MLVM_LIBS} TestArray PARENT_SCOPE)
