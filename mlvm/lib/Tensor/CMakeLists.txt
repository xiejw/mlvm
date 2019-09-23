add_library(Tensor
  Data.h
  Data.cpp
  Shape.h
  Shape.cpp
  Tensor.h
  Tensor.cpp
)

set(MLVM_LIBS ${MLVM_LIBS} Tensor PARENT_SCOPE)
