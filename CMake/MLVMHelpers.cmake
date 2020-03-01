# {{{1 Common Function
function(add_mlvm_executable NAME CPP_FILE)
  add_executable(${NAME} ${CPP_FILE})
  if(UNIX AND NOT APPLE)
    message("Build static binary for ${NAME}")
    target_link_libraries(${NAME} -static)  # Static binary.
  endif()
endfunction()
