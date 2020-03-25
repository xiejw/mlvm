# {{{1 Common Options
set(MLVM_COMMON_COMPILE_OPTIONS -Wall -Werror -Wextra)
set(MLVM_COMMON_INCLUDE_DIR ${CMAKE_CURRENT_SOURCE_DIR})

# {{{1 Common Function
function(add_mlvm_executable NAME CPP_FILE STATIC)
  add_executable(${NAME} ${CPP_FILE})
  if(${STATIC})
    message("Enable build static binary for ${NAME}")
    target_link_libraries(${NAME} -static)
  else()
    message("Disable build static binary for ${NAME}")
  endif()
endfunction()

function(add_mlvm_library)
  cmake_parse_arguments(
    MLVM_LIBRARY_PREFIX
    ""
    "NAME"
    "SRCS"
    ${ARGN}
  )

  # For name Foo
  #   - _NAME is MLVM_FOO
  #   - _LOWER_NAME is mlvm::foo
  set(_NAME MLVM_${MLVM_LIBRARY_PREFIX_NAME})
  string(TOUPPER ${_NAME} _NAME)
  string(TOLOWER mlvm::${MLVM_LIBRARY_PREFIX_NAME} _LOWER_NAME)

  add_library(${_NAME}
    ${MLVM_LIBRARY_PREFIX_SRCS}
  )

  add_library(${_LOWER_NAME}
    ALIAS ${_NAME}
  )

  target_include_directories(${_NAME} PUBLIC
    $<BUILD_INTERFACE:${MLVM_COMMON_INCLUDE_DIR}>
  )

  target_compile_options(${_NAME} PUBLIC
    ${MLVM_COMMON_COMPILE_OPTIONS})

endfunction()
