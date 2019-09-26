add_library(Comp
  Computation.h
  Computation.cpp
  Instruction.h
  Instruction.cpp
  OpType.h
  OpType.cpp
)

set(MLVM_LIBS ${MLVM_LIBS} Comp PARENT_SCOPE)
