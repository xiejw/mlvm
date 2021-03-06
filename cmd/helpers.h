// Helper methods for cmd/main.go

#define NE(e) _NO_ERR_IMPL_(e, __FILE__, __LINE__)

#define _NO_ERR_IMPL_(e, f, l)                                        \
        if (e) {                                                      \
                err = e;                                              \
                errDump("failed to exec op @ file %s line %d", f, l); \
                goto cleanup;                                         \
        }

#define S_PRINTF(prefix, t, suffix) \
        sdsCatPrintf(&s, prefix);   \
        vmTensorDump(&s, vm, t);    \
        sdsCatPrintf(&s, suffix);
