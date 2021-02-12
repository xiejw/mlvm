#include "testing/testing.h"

#include <string.h>

#include "opcode.h"
#include "opdefs.h"

// -----------------------------------------------------------------------------
// macros.
// -----------------------------------------------------------------------------

#define ASSERT_OPCODE_STR_AND_OPERAND_COUNT(opcode, num)                     \
        do {                                                                 \
                ASSERT_TRUE("OpCode def for " #opcode " mismatches",         \
                            0 == strcmp((opLookup(opcode, &def), def->name), \
                                        #opcode));                           \
                ASSERT_TRUE("OpCode def for " #opcode " too long",           \
                            opMaxCountOfChar >= strlen(def->name));          \
                ASSERT_TRUE("OpCode num for " #opcode " mismatches",         \
                            def->num_operands == num);                       \
                count++;                                                     \
        } while (0)

// -----------------------------------------------------------------------------
// unit tests.
// -----------------------------------------------------------------------------

static char* test_opcode_count()
{
        ASSERT_TRUE("total count for opcode. must <= byte (code_t)",
                    opTotalCount <= 255);
        return NULL;
}

static char* test_opcode_strs()
{
        struct opdef_t* def;        // used by macro.
        int             count = 0;  // used by macro.
        ASSERT_OPCODE_STR_AND_OPERAND_COUNT(OP_HALT, 0);
        ASSERT_OPCODE_STR_AND_OPERAND_COUNT(OP_PUSHBYTE, 1);
        ASSERT_OPCODE_STR_AND_OPERAND_COUNT(OP_LOADGLOBAL, 0);
        ASSERT_OPCODE_STR_AND_OPERAND_COUNT(OP_RETURN, 1);
        ASSERT_TRUE("not check all opcodes", count == opTotalCount);
        return NULL;
}

static char* test_make_op()
{
        vec_t(code_t) v = vecNew();
        size_t  offset  = 0;
        error_t err;

        // no_operand
        err = opMake(&v, OP_HALT);
        offset += 1;
        ASSERT_TRUE("no error", err == 0);
        ASSERT_TRUE("size", vecSize(v) == offset);
        ASSERT_TRUE("op code", v[offset - 1] == OP_HALT);

        // one operand uint8.
        err = opMake(&v, OP_PUSHBYTE, 123);
        offset += 2;
        ASSERT_TRUE("no error", err == 0);
        ASSERT_TRUE("size", vecSize(v) == offset);
        ASSERT_TRUE("op code", v[offset - 2] == OP_PUSHBYTE);
        ASSERT_TRUE("operand l", v[offset - 1] == 123);
        vecFree(v);
        return NULL;
}

static char* test_op_dump()
{
        error_t err;
        vec_t(code_t) v = vecNew();
        sds_t s         = sdsEmpty();

        err = opMake(&v, OP_PUSHBYTE, 123);
        ASSERT_TRUE("no error", err == OK);

        err = opMake(&v, OP_HALT);
        ASSERT_TRUE("no error", err == OK);

        err = opDump(&s, v, vecSize(v), /*prefix=*/"->");
        ASSERT_TRUE("no error", err == OK);

        char* expected_str = "->OP_PUSHBYTE    	123\n->OP_HALT        \n";
        ASSERT_TRUE("dump match", 0 == strcmp(s, expected_str) ||
                                      (printf("\n\ndump:\n%s\n", s), 0));

        sdsFree(s);
        vecFree(v);
        return NULL;
}

char* run_vm_opcode_suite()
{
        RUN_TEST(test_opcode_count);
        RUN_TEST(test_opcode_strs);
        RUN_TEST(test_make_op);
        RUN_TEST(test_op_dump);
        return NULL;
}
