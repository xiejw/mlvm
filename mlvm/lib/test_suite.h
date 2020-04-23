#include "mlvm/testing/test.h"

/* list_test.c */
extern char* run_list_test();

char* run_lib_suite() {
  RUN_SUITE(run_list_test);
  return NULL;
}
