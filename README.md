# Design

* Simple
* Readable
* Coherent

# Structure

## C Namespace

| Folder          | Namespace   |
| --------------- | ----------- |
| sprng           | `sprng_`    |
| ir/tensor       | `tensor_`   |
| runtime/kernel  | `kernel_`   |

## Lib

[`lib` folder](mlvm/lib) provides the basic data structure and algorithrms,
e.g., list, map, etc.

## Testing

[`testing` folder](mlvm/testing) provides a quite simple testing framework.
Check
- [`mlvm/testing/test.h`](mlvm/testing/test.h) for the macors and core ideas.
- [`cmd/test/main.c`](cmd/test/main.c) for the testing hierarchy.

