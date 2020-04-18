# vim: foldenable foldmethod=marker foldlevel=2
cmake_minimum_required(VERSION 3.9)

# {{{2 Configurations.
#
# {{{3 Compiler Part.
#
set(CMAKE_C_STANDARD 90)
set(CMAKE_C_COMPILER clang)

# {{{3 Common Helpers
#
list(APPEND CMAKE_MODULE_PATH ${CMAKE_CURRENT_LIST_DIR}/cmake)
include(MLVMHelpers)

# {{{2 Project
#
# {{{3 Set Project Name.
#
# Set after compiler before project.
#
# According to the CMake FAQ, Section "How do I use a different compiler",
# compiler change must be done before langguage is set, as quoted here:
#
#   Set the appropriate CMAKE_FOO_COMPILER variable(s) to a valid compiler name
#   or full path in a list file using set(). This must be done before any
#   language is set(ie before any project() or enable_language() command).
#
project(mlvm)

# {{{2 Dependencies.
#
# Putting supporting dependencies in bottom.
add_subdirectory(mlvm/random)

# {{{2 Executables.
#
# {{{3 Binary
set(BINARY prng)

option(MLVM_BUILD_STATIC "If ON, builds static binary" OFF)

add_mlvm_executable(${BINARY} cmd/prng/main.c ${MLVM_BUILD_STATIC})
target_link_libraries(${BINARY} mlvm::random)

