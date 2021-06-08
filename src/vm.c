#include "vm.h"
#include "vm_internal.h"

#include <stdlib.h>  // calloc
#include <string.h>  // memset

// mlvm
#include "primitives.h"

struct vm_t*
vmNew()
{
        return calloc(1, sizeof(struct vm_t));
}

void
vmFree(struct vm_t* vm)
{
        for (int i = 0; i < MLVM_MAX_TENSOR_COUNT; i++) {
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

struct shape_t*
vmShapeNew(struct vm_t* vm, int rank, int dims[])
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

error_t
vmBatch(struct vm_t* vm, size_t size, const struct oparg_t* args)
{
        const struct opopt_t* opt;
        const struct oparg_t* arg;
        error_t               err;

        for (size_t i = 0; i < size; i++) {
                arg = &args[i];
                if (arg->has_opt)
                        opt = &arg->opt;
                else {
                        opt = NULL;
                        assert(arg->opt.mode == 0);
                }

                err = vmExec(vm, arg->op, opt, arg->dst, arg->t1, arg->t2);
                if (err) {
                        return errEmitNote("failed to exec %d-th instruction.",
                                           i);
                }
        }

        return OK;
}

error_t
vmExec(struct vm_t* vm, enum opcode_t op, const struct opopt_t* opt, int dst,
       int lhs, int rhs)
{
        struct tensor_t* td = vmGrabHandle(vm, dst);
        struct tensor_t* t1 = NULL;
        struct tensor_t* t2 = NULL;

        if (lhs != -1) t1 = vmGrabHandle(vm, lhs);
        if (rhs != -1) t2 = vmGrabHandle(vm, rhs);

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
                                assert(OPT_MODE_GET_F_BIT(*opt));            \
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
                CASE_ELEWISE_OP(DIVIDE, Divide)
                CASE_ELEWISE_OP(MAX, Max)
                CASE_ELEWISE_OP(CMPL, CmpL)
                CASE_ELEWISE_OP(EQ, Eq)

#undef CASE_ELEWISE_OP

        case OP_MATMUL:
                assert(t1 != NULL);
                assert(t2 != NULL);
                if (td->dtype == F32) {
                        int trans_lhs = 0;
                        int trans_rhs = 0;
                        if (opt != NULL && opt->mode != OPT_MATMUL_TRANS_NOT) {
                                if (opt->mode == OPT_MATMUL_TRANS_LHS) {
                                        trans_lhs = 1;
                                } else {
                                        assert(opt->mode ==
                                               OPT_MATMUL_TRANS_RHS);
                                        trans_rhs = 1;
                                }
                        }
                        return vmOpMatmulF32(td, t1, t2, trans_lhs, trans_rhs);
                }

                return errNewWithNote(
                    ENOTIMPL, "unimpl for OP_MATMUL with dtype %d", td->dtype);

        case OP_RNG:
                assert(opt != NULL);
                assert(t1 == NULL);
                assert(t2 == NULL);
                assert(OPT_MODE_GET_R_BIT(*opt));
                assert((opt->mode & OPT_MODE_UNMASK) == 0);
                struct rng64_t rng = opt->r;
                if (td->dtype == F32) {
                        return vmOpRngF32(td, opt->mode & OPT_MODE_UNMASK,
                                          &rng);
                }

                return errNewWithNote(
                    ENOTIMPL, "unimpl for OP_RNG with dtype %d", td->dtype);

        case OP_ARGMAX:
                assert(opt == NULL);
                assert(t1 != NULL);
                assert(t2 == NULL);

                if (td->dtype == F32) {
                        return vmOpArgMaxF32(td, t1);
                }

                return errNewWithNote(
                    ENOTIMPL, "unimpl for OP_ARGMAX with dtype %d", td->dtype);

        case OP_REDUCE:
                assert(opt != NULL);
                assert((opt->mode & OPT_MODE_UNMASK) == 0);
                assert(t1 != NULL);
                assert(t2 == NULL);
                if (td->dtype != F32) {
                        return errNewWithNote(
                            ENOTIMPL, "unimpl for OP_REDUCE with dtype %d",
                            td->dtype);
                }

                int axis = 0;
                if (OPT_MODE_GET_I_BIT(*opt)) {
                        axis = opt->i;
                } else {
                        assert(opt->i == 0);
                }
                return vmOpReduceF32(td, t1, opt->mode & OPT_MODE_UNMASK, axis);

        case OP_FILL:
                assert(t1 == NULL);
                assert(t2 == NULL);
                if (td->dtype != F32) {
                        return errNewWithNote(
                            ENOTIMPL, "unimpl for OP_FILL with dtype %d",
                            td->dtype);
                }

                if (opt == NULL) {
                        memset(td->data, 0,
                               td->shape->size * sizeof(float32_t));
                        return OK;
                }
                assert((opt->mode & OPT_MODE_UNMASK) == 0);
                assert(OPT_MODE_GET_F_BIT(*opt));
                return vmOpFillF32(td, opt->f);

        case OP_ISQRT:
                assert(t2 == NULL);
                if (td->dtype != F32) {
                        return errNewWithNote(
                            ENOTIMPL, "unimpl for OP_ISQRT with dtype %d",
                            td->dtype);
                }

                if (opt == NULL) {
                        return vmOpISqrtF32(td, t1, NULL, /*mode=*/-1);
                }

                assert((opt->mode & OPT_MODE_UNMASK) >> 1 == 0);
                assert(OPT_MODE_GET_F_BIT(*opt));
                return vmOpISqrtF32(td, t1, &opt->f,
                                    /*mode=*/opt->mode & OPT_MODE_UNMASK);

        case OP_LS_SCEL:
                assert(t1 != NULL);
                assert(t2 != NULL);
                struct tensor_t* tg = NULL;
                if (td->dtype != F32) {
                        return errNewWithNote(
                            ENOTIMPL, "unimpl for OP_LS_SCEL with dtype %d",
                            td->dtype);
                }

                if (opt != NULL) {
                        if (OPT_MODE_GET_I_BIT(*opt)) {
                                tg = vmGrabHandle(
                                    vm,
                                    opt->i);  // handle of the gradient.
                        } else {
                                assert(opt->i == 0);
                        }
                }
                return vmOpLossSCELF32(td, t1, t2, tg);

        default:
                return errNewWithNote(ENOTIMPL, "unimpl op for vmExec: %d", op);
        }
}
