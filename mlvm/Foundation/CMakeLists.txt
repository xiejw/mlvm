add_mlvm_library(
  NAME
    foundation

  SRCS
    StatusOr.h
    Status.h
    Status.cc
    Logging.h
    Logging.cc
    Utilities.h
    Utilities.cc

  TESTS
    StatusTest.cc
    StatusOrTest.cc

  DEPS
    absl::strings
)

