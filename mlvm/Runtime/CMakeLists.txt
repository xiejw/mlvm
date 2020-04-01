add_mlvm_library(
  NAME
    rt

  SRCS
    Evaluator.h
    Evaluator.cc

  DEPS
    mlvm::ir
    mlvm::foundation
)
