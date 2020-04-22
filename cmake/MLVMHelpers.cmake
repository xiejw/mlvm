# {{{1 Common Options
set(MLVM_COMMON_COMPILE_OPTIONS -Wall -Werror -Wextra)
set(MLVM_COMMON_INCLUDE_DIR ${CMAKE_CURRENT_SOURCE_DIR})

# {{{1 Common Function
function(add_mlvm_executable NAME CPP_FILE STATIC)
  add_executable(${NAME} ${CPP_FILE})
  if(${STATIC})
    message("## Enable build static binary for ${NAME}")
    target_link_libraries(${NAME} -static)
  else()
    message("## Disable build static binary for ${NAME}")
  endif()
endfunction()

function(add_mlvm_library)
  cmake_parse_arguments(
    MLVM_LIBRARY_PREFIX
    ""
    "NAME"
    "SRCS;TESTS;DEPS"
    ${ARGN}
  )

  # Set up some names used by later. For example, if the NAME is Foo, then
  #
  #   - `_PUBLIC_NAME` is `mlvm::foo`. This is user-facing name.
  #   - `_INTERNAL_NAME` is `MLVM_FOO`. This is mainly used for internal target.
  #
  string(TOLOWER mlvm::${MLVM_LIBRARY_PREFIX_NAME} _PUBLIC_NAME)
  string(TOUPPER MLVM_${MLVM_LIBRARY_PREFIX_NAME} _INTERNAL_NAME)

  ##############################################################################
  # Library
  ##############################################################################
  add_library(${_INTERNAL_NAME}
    ${MLVM_LIBRARY_PREFIX_SRCS}
  )

  target_include_directories(${_INTERNAL_NAME} PUBLIC
    $<BUILD_INTERFACE:${MLVM_COMMON_INCLUDE_DIR}>
  )

  target_compile_options(${_INTERNAL_NAME} PUBLIC
    ${MLVM_COMMON_COMPILE_OPTIONS})

  target_link_libraries(${_INTERNAL_NAME} PUBLIC ${MLVM_LIBRARY_PREFIX_DEPS})

  ##############################################################################
  # Alias
  ##############################################################################

  add_library(${_PUBLIC_NAME} ALIAS ${_INTERNAL_NAME})

  ##############################################################################
  # Test
  ##############################################################################
  if (MLVM_LIBRARY_PREFIX_TESTS)
    set(_TEST_NAME TEST_${_INTERNAL_NAME})

    # Test cases should be linked in cmake as OBJECT library.
    add_library(${_TEST_NAME} OBJECT ${MLVM_LIBRARY_PREFIX_TESTS})
    target_link_libraries(${_TEST_NAME} PUBLIC ${_INTERNAL_NAME})
    target_link_libraries(${_TEST_NAME} PUBLIC mlvm::testing)

    # Added an alias.
    add_library(test::${_PUBLIC_NAME} ALIAS ${_TEST_NAME})
  endif()
endfunction()

function(target_link_mlvm_test_libraries NAME LIB)
  target_link_libraries(${NAME} test::${LIB})
endfunction()
