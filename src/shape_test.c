#include "testing/testing.h"

#include "vm.h"

static char*
test_shape_init()
{
        struct shape_t* s = spNew(2, (int[]){2, 3});
        spDecRef(s);
        return NULL;
}

char*
run_shape_suite()
{
        RUN_TEST(test_shape_init);
        return NULL;
}
