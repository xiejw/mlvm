add_mlvm_library(
  NAME
    foundation

  SRCS
    Status.h
    StatusOr.h
    Status.cc
    Utilities.h
    Utilities.cc

  TESTS
    StatusTest.cc
    StatusOrTest.cc
)

