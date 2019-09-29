add_library(Array
  Data.h
  Data.cpp
  Shape.h
  Shape.cpp
  Array.h
  Array.cpp
)

set(MLVM_LIBS ${MLVM_LIBS} Array PARENT_SCOPE)
