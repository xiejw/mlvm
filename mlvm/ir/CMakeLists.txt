add_mlvm_library(
  NAME
    ir

  SRCS
    ir.h
    tensor.h
    tensor.c

    context.h
    context.c
    function.h
    operand.c
    instruction.c
    function.c

    error.h
    macros.h

  TESTS
    test_suite.h
    context_test.c
)
