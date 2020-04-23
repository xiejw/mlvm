# Design

* Simple
* Readable
* Coherent

# Structure

## C Namespace

| Folder          | Namespace   | CMake Target   |
| --------------- | ----------- | -------------- |
| sprng           | `sprng_`    | `mlvm::sprng`  |
| ir/tensor       | `tensor_`   | `mlvm::ir`     |
| runtime/kernel  | `kernel_`   | `mlvm::kernel` |

## Lib

[`lib` folder](mlvm/lib) provides the basic data structure and algorithrms,
e.g., list, map, etc.

The cmake target is `mlvm::lib`.

## Testing

[`testing` folder](mlvm/testing) provides a quite simple testing framework.
Check
- [`mlvm/testing/test.h`](mlvm/testing/test.h) for the macors and core ideas.
- [`cmd/test/main.c`](cmd/test/main.c) for the testing hierarchy.

The cmake target is `mlvm::testing`.
