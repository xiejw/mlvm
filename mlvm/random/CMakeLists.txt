add_mlvm_library(
  NAME
    random

  SRCS
    normal.h
    normal.c
    sprng64.h
    sprng64.c

  DEPS
    -lm
)
