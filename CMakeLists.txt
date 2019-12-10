# vim: foldenable foldmethod=marker foldlevel=1
cmake_minimum_required(VERSION 3.5)
project(mlvm)

# {{{2 Configurations.

# {{{3 Compiler Part.
# Set the C++17 standard and Clang for compiler.
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_C_COMPILER clang)
set(CMAKE_CXX_COMPILER clang++)

# Adjust some options.
add_compile_options(-Wall)

# {{{3 Common Libraries, like gflags, gtest.
include(${CMAKE_CURRENT_SOURCE_DIR}/../common-lib/CommonLib.cmake)

# {{{2 Includes.
# Add local directory into -I
include_directories(.)

# {{{2 Dependencies.
#
# Putting supporting dependencies in bottom.
add_subdirectory (mlvm/Local)

# {{{2 Executables.
# {{{3 Main
new_executable(example cmd/main.cpp)

target_link_libraries(example ${MLVM_LIBS})

