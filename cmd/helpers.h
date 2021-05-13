// Helper methods for cmd/main.go

#define NE(e) NO_ERR_IMPL_(e, __FILE__, __LINE__)

#define NO_ERR_IMPL_(e, f, l)                                           \
        if (e) {                                                        \
                err = e;                                                \
                errDump("failed to exec op @ file %s line %d\n", f, l); \
                goto cleanup;                                           \
        }

#define R1S(vm, s1)     vmShapeNew(vm, 1, (int[]){(s1)});
#define R2S(vm, s1, s2) vmShapeNew(vm, 2, (int[]){(s1), (s2)});

#define S_PRINTF(prefix, t, suffix) \
        sdsCatPrintf(&s, prefix);   \
        vmTensorDump(&s, vm, t);    \
        sdsCatPrintf(&s, suffix);
