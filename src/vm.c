#include "vm.h"

#include "primitives.h"
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
        struct list_t* cur = vm->shapes;
        struct list_t* nxt;
        while (cur != NULL) {
                nxt = cur->next;
                spDecRef(cur->data);
                free(cur);
                cur = nxt;
        }
        free(vm);
}

struct shape_t* vmShapeNew(struct vm_t* vm, int rank, int dims[])
{
        struct shape_t* s = spNew(rank, dims);

        struct list_t* n = malloc(sizeof(struct list_t));
        n->data          = s;
        if (vm->shapes == NULL) {
                n->next    = NULL;
                vm->shapes = n;
        } else {
                n->next    = vm->shapes;
                vm->shapes = n;
        }

        return s;
}

error_t vmTensorSwap(struct vm_t* vm, int t, _mut_ void** data)
{
        struct tensor_t* ts  = vmGrabHandle(vm, t);
        void*            old = ts->data;
        ts->data             = *data;
        *data                = old;
        return OK;
}

error_t vmExec(struct vm_t* vm, enum opcode_t op, const struct opopt_t* opt,
               int dst, int lhs, int rhs)
{
        struct tensor_t* td = vmGrabHandle(vm, dst);
        struct tensor_t* t1 = NULL;
        struct tensor_t* t2 = NULL;

        if (lhs != VM_UNUSED) t1 = vmGrabHandle(vm, lhs);
        if (rhs != VM_UNUSED) t2 = vmGrabHandle(vm, rhs);

        switch (op) {
#define CASE_ELEWISE_OP(OP, API)                                             \
        case OP_##OP:                                                        \
                assert(t1 != NULL);                                          \
                if (td->dtype == F32) {                                      \
                        if (opt == NULL) {                                   \
                                assert(t2 != NULL);                          \
                                return vmOp##API##F32(td, t1, t2);           \
                        } else {                                             \
                                assert(t2 == NULL);                          \
                                return vmOp##API##SF32(td, t1, opt->f);      \
                        }                                                    \
                }                                                            \
                                                                             \
                return errNewWithNote(ENOTIMPL,                              \
                                      "unimpl for OP_" #OP " with dtype %d", \
                                      td->dtype);

                CASE_ELEWISE_OP(ADD, Add)
                CASE_ELEWISE_OP(MUL, Mul)
                CASE_ELEWISE_OP(MINUS, Minus)
                CASE_ELEWISE_OP(MAX, Max)
                CASE_ELEWISE_OP(CMPL, CmpL)

#undef CASE_ELEWISE_OP

        case OP_MATMUL:
                assert(t1 != NULL);
                assert(t2 != NULL);
                assert(opt == NULL);
                if (td->dtype == F32) {
                        return vmOpMatmulF32(td, t1, t2, 0, 0);
                }

                return errNewWithNote(
                    ENOTIMPL, "unimpl for OP_MATMUL with dtype %d", td->dtype);

        case OP_RNG:
                assert(opt != NULL);
                assert(opt->mode == 0);
                assert(t1 == NULL);
                assert(t2 == NULL);
                if (td->dtype == F32) {
                        return vmOpRngF32(td, opt->mode, opt->rng_seed);
                }

                return errNewWithNote(
                    ENOTIMPL, "unimpl for OP_RNG with dtype %d", td->dtype);

        case OP_REDUCE:
                assert(opt != NULL);
                assert(opt->mode == 0);
                assert(t1 != NULL);
                assert(t2 == NULL);
                if (td->dtype == F32) {
                        return vmOpReduceF32(td, t1, opt->mode);
                }

                return errNewWithNote(
                    ENOTIMPL, "unimpl for OP_REDUCE with dtype %d", td->dtype);

        case OP_LS_SCEL:
                assert(t1 != NULL);
                assert(t2 != NULL);
                assert(opt == NULL);
                if (td->dtype == F32) {
                        return vmOpLossSCELF32(td, t1, t2);
                }

                return errNewWithNote(
                    ENOTIMPL, "unimpl for OP_LS_SCEL with dtype %d", td->dtype);

        default:
                return errNewWithNote(ENOTIMPL, "unimpl op for vmExec: %d", op);
        }
}
