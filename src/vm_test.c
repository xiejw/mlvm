#include "testing/testing.h"

#include "vm.h"

#include <string.h>

static char*
test_vm_new()
{
        struct vm_t* vm = vmNew();
        ASSERT_TRUE("vm", vm != NULL);
        vmFree(vm);
        return NULL;
}

char*
run_vm_suite()
{
        RUN_TEST(test_vm_new);
        return NULL;
}
