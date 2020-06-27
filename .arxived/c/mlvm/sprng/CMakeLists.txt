add_mlvm_library(
  NAME
    sprng

  SRCS
    sprng.h
    sprng_api.h
    sprng64.h
    sprng64.c
    normal.h
    normal.c

  TESTS
    test_suite.h
    normal_test.c

  DEPS
    -lm
)
