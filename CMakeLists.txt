# vim: foldenable foldmethod=marker foldlevel=1
cmake_minimum_required(VERSION 3.5)

# {{{2 Configurations.
#
# {{{3 Compiler Part.
#
set(CMAKE_CXX_STANDARD 20)
set(CMAKE_C_COMPILER clang)
set(CMAKE_CXX_COMPILER clang++)

# {{{3 Adjust some options for compiler globally.
#
add_compile_options(-Wall)

# {{{2 Project
#
# Set after compiler to avoid infinite loop in absl.
project(mlvm)

# {{{3 Common Helpers
#
list(APPEND CMAKE_MODULE_PATH ${CMAKE_CURRENT_LIST_DIR}/CMake)
include(MLVMHelpers)

# {{{2 Dependencies.
#
# Putting supporting dependencies in bottom.
add_subdirectory(libraries/absl)

# {{{2 Executables.
#
# {{{3 Binary compile
#
set(BINARY compile)
add_mlvm_executable(${BINARY} cmd/compile/main.cpp)
target_link_libraries(${BINARY} absl::strings)
