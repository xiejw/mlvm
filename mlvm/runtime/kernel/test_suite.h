/* matmul_test.c */
extern char* run_matmul_test();

char* run_kernel_suite() {
  RUN_SUITE(run_matmul_test);
  return NULL;
}
