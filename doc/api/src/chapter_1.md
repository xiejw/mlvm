# MLVM

TL;DR; MLVM is a fast VM to execute machine learning pritimives.

The MLVM is a VM I want to, and will, use for my own projects for next decades.
It might not fit others' needs; but I can understand and reason about it. It is
simple and efficient. It does not do any checks&ndahs;behave like a machine. And
it provides raw accesses to the underlying stack, so it is possible to do
in-place operations and swap pointers, which could be dangerous but certainly
are critical for performance. Memory management is in user's hands, and
auto-grad is not integrated&ndash;yet it is just simple and fast.

## Design

The first level of the VM is designed to be a system library, not compiler tool.
It is super simple, easy to understand, and super efficient. Toward efficiency,
there will be no check, just like a ordinary machine. So caller should be very
confident about its correctness. A programming language and compiler built on
top of the VM is highly recommended, but not required.

Another design point is: How to define the model for parallelism. In particular,
how to use the multi-core system to gain performance. Two ways are possible. The
first one is offloading the computation to co-processors. With that, we need to
introduce async value, multi-threads for heavy operations e.g., matmuls. The
known challenges are also obvious: How to avoid cache line miss during processor
switch, and how to design a scheduling system which handles cheap operations
well&ndash;a chain of cheap operations might put most of the processors idle.

The alternative is to design the VM for single processor execution only, and do
it very well. Caller can launch multiple VM's each for single processor and
adjust the affinity if needed. With that, MPI style collective are needed. For
CPU within same OS process, we can provide a shared memory implementation for
the collective to avoid data copying.
