add_library(MLVM_COMPUTATION
  Tensor.h
  Tensor.cc
)

target_include_directories(MLVM_COMPUTATION PUBLIC
  $<BUILD_INTERFACE:${CMAKE_SOURCE_DIR}>
)
