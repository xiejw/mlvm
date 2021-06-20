# Design

The design goals of MLVM is
* Simple
* Readable
* Coherent

It has only one concept, the `struct vm_t` (in [vm.h](include/vm.h)). It
supports eager execution, in-place mutation, and batch mode.

Check [Blueberry](https://github.com/xiejw/blueberry) for one possible way to
use it.

## Compile Flags

```
MLVM_TENSOR_DUMP_SIZE    The number of elements to print for Tensor, e.g.,
                         CFLAGS+=-DMLVM_TENSOR_DUMP_SIZE=3 make -B regression
MLVM_MAX_TENSOR_COUNT    The number of tensors VM supports, defaults to 128.
```

## Blis

Compile blis in `../blis`

```
# cd blis
$ git clone https://github.com/flame/blis.git
$ ./configure auto
$ make
```

Compile MLVM with blis, e.g.,

```
# cd mlvm
$ make BLIS=1 RELESAE=1 -B
$ make BLIS=1 RELESAE=1 libmlvm
$ make BLIS=1 test -B
```
