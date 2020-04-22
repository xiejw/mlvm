add_mlvm_library(
  NAME
    random

  SRCS
    random.h
    normal.h
    normal.c
    sprng64.h
    sprng64.c

  TESTS
    random_test.h
    normal_test.c

  DEPS
    -lm
)
