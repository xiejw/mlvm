add_library(Local
  Data.h
  Data.cpp
)

set(MLVM_LIBS ${MLVM_LIBS} Local PARENT_SCOPE)
