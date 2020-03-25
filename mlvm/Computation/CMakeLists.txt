# add_library(MLVM_COMPUTATION
#   Tensor.h
#   Tensor.cc
# )
# 
# target_include_directories(MLVM_COMPUTATION PUBLIC
#   $<BUILD_INTERFACE:${MLVM_COMMON_INCLUDE_DIR}>
# )
# 
# target_compile_options(MLVM_COMPUTATION PUBLIC ${MLVM_COMMON_COMPILE_OPTIONS})

add_mlvm_library(
  NAME
    Computation

  SRCS
    Tensor.h
    Tensor.cc
)
