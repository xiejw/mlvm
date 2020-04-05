add_mlvm_library(
  NAME
    op

  SRCS
    Einsum.h
    Einsum.cc

  TESTS
    EinsumTest.cc

  DEPS
    absl::strings
    absl::flat_hash_map
)
