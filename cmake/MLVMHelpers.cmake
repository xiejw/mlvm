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

set(MLVM_COMMON_COMPILE_OPTIONS -Wall -Werror PARENT_SCOPE)
