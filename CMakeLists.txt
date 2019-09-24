cmake_minimum_required(VERSION 3.13)
project(mlvm)

# Common Libraries
include(${CMAKE_CURRENT_SOURCE_DIR}/../common-lib/CommonLib.cmake)

# Set the C++14 standard
set(CMAKE_CXX_STANDARD 14)
set(CMAKE_C_COMPILER clang)
set(CMAKE_CXX_COMPILER clang++)
add_compile_options(-Wall)

# I../includes
include_directories(.)
include_directories(${GFLAGS_INCLUDES})

# All libraries for MLVM. Due to static libraries linking order requiment, put
# supporting libraries in end.
add_subdirectory (mlvm/lib/Tensor)
add_subdirectory (mlvm/lib/Support)

add_executable(test_app examples/main.cpp)
if(UNIX AND NOT APPLE)
  message("Build static binary for test_app")
  target_link_libraries(test_app -static)  # Static binary.
endif()

target_link_libraries(test_app ${MLVM_LIBS})
target_link_libraries(test_app ${GFLAGS_LIBS})

# Tests
add_executable(tests mlvm/test/main.cpp)
include(mlvm/test/Test.cmake)  # Contains all test folders.

target_link_libraries(tests ${TEST_LIBS})
target_link_libraries(tests ${MLVM_LIBS})

target_include_directories(tests PUBLIC ${GTEST_INCLUDES})
target_link_libraries(tests ${GTEST_LIBS})
