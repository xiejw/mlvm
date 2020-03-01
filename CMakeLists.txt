# vim: foldenable foldmethod=marker foldlevel=1

cmake_minimum_required(VERSION 3.5)

# {{{2 Configurations.

# {{{3 Compiler Part.
# Set the C++20 standard and Clang for compiler.
set(CMAKE_CXX_STANDARD 20)
set(CMAKE_C_COMPILER clang)
set(CMAKE_CXX_COMPILER clang++)

# {{{3 Adjust some options for compiler globally.
add_compile_options(-Wall)

# {{{2 Project
# Set after compiler to avoid infinite loop in absl.
project(mlvm)

# {{{3 Common Libraries, like gflags, gtest.
include(${CMAKE_CURRENT_SOURCE_DIR}/../common-lib/CommonLib.cmake)

# {{{2 Dependencies.
# Putting supporting dependencies in bottom.
add_subdirectory(libraries/absl)
# add_subdirectory (eva/Cron)

# {{{2 Executables.
# {{{3 Compile
new_executable(compile cmd/compile/main.cpp)

target_link_libraries(compile absl::strings)


