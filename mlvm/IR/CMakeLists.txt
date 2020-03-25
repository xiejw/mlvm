add_mlvm_library(
  NAME
    ir

  SRCS
    Tensor.h
    Tensor.cc
    Instruction.h
    Instruction.cc
    Function.h
    Function.cc

  DEPS

    absl::strings
)
