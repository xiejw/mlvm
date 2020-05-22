/* matmul_test.c */
extern char* run_matmul_test();

/* mul_test.c */
extern char* run_mul_test();

/* add_test.c */
extern char* run_add_test();

char* run_kernel_suite() {
  RUN_SUITE(run_matmul_test);
  RUN_SUITE(run_mul_test);
  RUN_SUITE(run_add_test);
  return NULL;
}
