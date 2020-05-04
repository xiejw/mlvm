/* context_test.c */
extern char* run_context_test();

char* run_ir_suite() {
  RUN_SUITE(run_context_test);
  return NULL;
}
