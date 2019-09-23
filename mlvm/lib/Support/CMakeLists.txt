add_library(Support
  Error.h
  Error.cpp
)

set(MLVM_LIBS ${MLVM_LIBS} Support PARENT_SCOPE)
