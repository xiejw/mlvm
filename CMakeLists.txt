# vim: foldenable foldmethod=marker foldlevel=2
cmake_minimum_required(VERSION 3.9)

# {{{2 Configurations.
#
# {{{3 Compiler Part.
#
set(CMAKE_CXX_STANDARD 20)
set(CMAKE_C_COMPILER clang)
set(CMAKE_CXX_COMPILER clang++)

# {{{3 Common Helpers
#
list(APPEND CMAKE_MODULE_PATH ${CMAKE_CURRENT_LIST_DIR}/cmake)
include(MLVMHelpers)

# {{{2 Project
#
# Set after compiler to avoid infinite loop in absl.
project(mlvm)

# {{{2 Dependencies.
#
# Putting supporting dependencies in bottom.
add_subdirectory(.libraries/absl)
add_subdirectory(mlvm/Computation)

# {{{2 Executables.
#
# {{{3 Binary demo
#
set(BINARY demo)

option(MLVM_BUILD_STATIC "If ON, builds static binary" OFF)

add_mlvm_executable(${BINARY} cmd/demo/main.cc ${MLVM_BUILD_STATIC})
target_link_libraries(${BINARY} MLVM_COMPUTATION)
target_link_libraries(${BINARY} absl::strings)
target_compile_options(${BINARY} PUBLIC ${MLVM_COMMON_COMPILE_OPTIONS})
