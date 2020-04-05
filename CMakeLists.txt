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
add_subdirectory(mlvm/IR)
add_subdirectory(mlvm/Op)
add_subdirectory(mlvm/Runtime)
add_subdirectory(mlvm/Foundation)
add_subdirectory(.libraries/absl)
add_subdirectory(.libraries/gtest)

# {{{2 Executables.
#
# {{{3 Binary demo
#
set(BINARY demo)

option(MLVM_BUILD_STATIC "If ON, builds static binary" OFF)

add_mlvm_executable(${BINARY} cmd/demo/main.cc ${MLVM_BUILD_STATIC})
target_link_libraries(${BINARY} mlvm::ir)
target_link_libraries(${BINARY} mlvm::op)
target_link_libraries(${BINARY} mlvm::rt)
target_link_libraries(${BINARY} mlvm::foundation)
target_link_libraries(${BINARY} absl::strings)
target_compile_options(${BINARY} PUBLIC ${MLVM_COMMON_COMPILE_OPTIONS})

# {{{2 Tests.
#
# {{{3 Main Test.
#
if (NOT "${CMAKE_BUILD_TYPE}" STREQUAL "RELEASE")
  message("## Generate test for debug mode")
  add_executable(test cmd/test/main.cc)
  target_link_mlvm_test_libraries(test mlvm::foundation)
  target_link_libraries(test gtest )
else()
  message("## Skip test for release mode")
endif ()
