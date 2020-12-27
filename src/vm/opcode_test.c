#include "testing/testing.h"

#include "opcode.h"

static char* test_opcode_count() {
  ASSERT_TRUE("max count for opcode", OP_END <= 255);
  return NULL;
}

static char* test_opcode_strs() {
  opdeft* def;

#define ASSERT_OPCODE_STR_AND_OPERAND_COUNT(opcode, num)        \
  ASSERT_TRUE("OpCode def for " #opcode " mismatches",          \
              (opLookup(opcode, &def), def->name) == #opcode && \
                  def->num_operands == num);

  ASSERT_OPCODE_STR_AND_OPERAND_COUNT(OP_CONST, 1);
  ASSERT_OPCODE_STR_AND_OPERAND_COUNT(OP_POP, 0);

#undef ASSERT_OPCODE_STR_AND_OPERAND_COUNT
  return NULL;
}

char* test_make_op() {
  vect(codet) v = NULL;
  size_t offset = 0;

  // no_operand
  errort err = opMake(OP_POP, &v);
  offset += 1;
  ASSERT_TRUE("no error", err == 0);
  ASSERT_TRUE("size is 0", vecSize(v) == offset);
  ASSERT_TRUE("op code", v[offset - 1] == OP_POP);

  // one operand uint16.
  offset += 3;
  err = opMake(OP_CONST, &v, 123);
  ASSERT_TRUE("no error", err == 0);
  ASSERT_TRUE("size is 0", vecSize(v) == offset);
  ASSERT_TRUE("op code", v[offset - 3] == OP_CONST);
  ASSERT_TRUE("operand h", v[offset - 2] == 0);
  ASSERT_TRUE("operand l", v[offset - 1] == 123);
  vecFree(v);
  return NULL;
}

char* run_opcode_suite() {
  RUN_TEST(test_opcode_count);
  RUN_TEST(test_opcode_strs);
  RUN_TEST(test_make_op);
  return NULL;
}
