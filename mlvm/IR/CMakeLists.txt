add_mlvm_library(
  NAME
    ir

  SRCS
    Tensor.h
    Tensor.cc
    Instruction.h
    Instruction.cc

  DEPS

    absl::strings
)
