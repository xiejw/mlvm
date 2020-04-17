add_mlvm_library(
  NAME
    foundation

  SRCS
    StatusOr.h
    Status.h
    Status.cc
    Logger.h
    Logger.cc
    Logging.h
    Logging.cc
    Utilities.h
    Utilities.cc

  TESTS
    StatusTest.cc
    StatusOrTest.cc

  DEPS
    absl::flags
    absl::flags_parse # Required for Parsing CMD.
    absl::strings
)

