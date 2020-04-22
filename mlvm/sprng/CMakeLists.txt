add_mlvm_library(
  NAME
    sprng

  SRCS
    sprng.h
    sprng_interface.h
    sprng64.h
    sprng64.c
    normal.h
    normal.c

  TESTS
    sprng_test.h
    normal_test.c

  DEPS
    -lm
)
