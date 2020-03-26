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
add_subdirectory(mlvm/IR)
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
target_link_libraries(${BINARY} mlvm::foundation)
target_link_libraries(${BINARY} absl::strings)
target_compile_options(${BINARY} PUBLIC ${MLVM_COMMON_COMPILE_OPTIONS})

# {{{2 Tests.
if (NOT "${CMAKE_BUILD_TYPE}" STREQUAL "RELEASE")
  message("## Generate test for debug mode")
  add_executable(test cmd/test/main.cc)
  target_link_mlvm_test_libraries(test mlvm::foundation)
  target_link_libraries(test gtest )
else()
  message("## Skip test for release mode")
endif ()
