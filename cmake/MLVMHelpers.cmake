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

# {{{1 Common Options
set(MLVM_COMMON_COMPILE_OPTIONS -Wall -Werror -Wextra)
set(MLVM_COMMON_INCLUDE_DIR ${CMAKE_CURRENT_SOURCE_DIR})
