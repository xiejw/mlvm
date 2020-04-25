/* matmul_test.c */
extern char* run_matmul_test();

/* mul_test.c */
extern char* run_mul_test();

char* run_kernel_suite() {
  RUN_SUITE(run_matmul_test);
  RUN_SUITE(run_mul_test);
  return NULL;
}
