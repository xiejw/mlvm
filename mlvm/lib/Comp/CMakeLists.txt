add_library(Comp
  Computation.h
  Computation.cpp
)

set(MLVM_LIBS ${MLVM_LIBS} Comp PARENT_SCOPE)
