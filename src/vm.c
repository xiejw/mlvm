#include "vm.h"

#include "op.h"
#include "vm_internal.h"

#include <string.h>  // memset

struct vm_t* vmNew()
{
        size_t       size = sizeof(struct vm_t);
        struct vm_t* vm   = malloc(size);
        memset(vm, 0, size);
        return vm;
}

void vmFree(struct vm_t* vm)
{
        for (int i = 0; i < MAX_TENSOR_COUNT; i++) {
                struct tensor_t* t = &vm->handles[i];
                if (t->used) vmReleaseHandle(t);
        }
        free(vm);
}

error_t vmExec(struct vm_t* vm, enum opcode_t op, const struct opopt_t* opt,
               int dst, int lhs, int rhs)
{
        struct tensor_t* td = vmGrabHandle(vm, dst);
        struct tensor_t* t1 = NULL;
        struct tensor_t* t2 = NULL;

        if (lhs != -1) t1 = vmGrabHandle(vm, lhs);
        if (rhs != -1) t2 = vmGrabHandle(vm, rhs);

        switch (op) {
#define CASE_ELEWISE_OP(OP, API)                                             \
        case OP_##OP:                                                        \
                assert(opt == NULL);                                         \
                assert(t1 != NULL);                                          \
                assert(t2 != NULL);                                          \
                if (td->dtype == F32) {                                      \
                        return vmOp##API##F32(td, t1, t2);                   \
                }                                                            \
                                                                             \
                return errNewWithNote(ENOTIMPL,                              \
                                      "unimpl for OP_" #OP " with dtype %d", \
                                      td->dtype);

                CASE_ELEWISE_OP(ADD, Add)
                CASE_ELEWISE_OP(MUL, Mul)
                CASE_ELEWISE_OP(MINUS, Minus)

        case OP_RNG:
                assert(opt != NULL);
                assert(opt->mode == 0);
                assert(t1 == NULL);
                assert(t2 == NULL);
                if (td->dtype == F32) {
                        return vmOpcRngF32(td, opt->mode, opt->rng_seed);
                }

                return errNewWithNote(
                    ENOTIMPL, "unimpl for OP_RNG with dtype %d", td->dtype);

        default:
                return errNewWithNote(ENOTIMPL, "unimpl for vmExec");
        }
}

void vmSync() {}
