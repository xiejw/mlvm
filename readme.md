# Design

* Simple
* Readable
* Coherent

## Compile Flags

```
MLVM_TENSOR_DUMP_SIZE           The number of elements to print for Tensor.
    e.g. CFLAGS+=-DMLVM_TENSOR_DUMP_SIZE=3 make -B regression
```

## Blis

Compile blis in `../blis`

```
# cd blis
$ git clone https://github.com/flame/blis.git
$ ./configure auto
$ make
```

Obtain the config name
```
# cd blis
$ grep ^CONFIG_NAME config.mk
CONFIG_NAME       := haswell
```

Compile MLVM with blis, e.g.,

```
# cd mlvm
$ make BLIS=haswell RELESAE=1 -B
```
