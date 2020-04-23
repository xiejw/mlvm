add_mlvm_library(
  NAME
    kernel

  SRCS
    kernel.h
    add.c
    mul.c
    matmul.c
    shape_util.h
    shape_util.c

  TESTS
    test_suite.h
    matmul_test.c

  DEPS
    mlvm::ir
)
