#include "testing/testing.h"

#include "vm.h"

static char*
test_shape_init()
{
        struct shape_t* s = spNew(2, (int[]){3, 4});
        ASSERT_TRUE("rank", s->rank == 2);
        ASSERT_TRUE("dim 0", s->dims[0] == 3);
        ASSERT_TRUE("dim 1", s->dims[1] == 4);
        ASSERT_TRUE("size", s->size == 12);
        spDecRef(s);
        return NULL;
}

char*
run_vm_suite()
{
        RUN_TEST(test_shape_init);
        return NULL;
}
