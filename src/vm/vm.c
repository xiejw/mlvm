#include "vm.h"

error_t vmLaunch(struct vm_t* vm) { return OK; }

float vmComsumedSizeInMB(struct vm_t* vm)
{
        return (float)(((double)vm->size_used) / 1024 / 1024);
}
