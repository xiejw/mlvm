Topic: Pure Computation Graph
======================

- Draft: 2019-08
- State: Abandoned

History
-------

For a pure computation graph, computaiton graph dependency is defined by result usage.

In addition, given this is no side effect for any pure computation graph, a dead
node, whose results have no users, can be removed safely.

Variable
--------

Graph containing variable have barriers and side effects.

### Barrier

- Any variable read should occure before next write, but after previous write.
- Any variable write must follow prevous write and all readers in between.

### Side effects

- Any variable read has no side effects. So, it can be removed if it belongs to
  a dead branch.
- Any variable write has side effects. Variable write which have no variable
  read followed becomes dead branch.  But, we should introduce a sink node in
  the graph and adds all variables writes as control edges to it.

### Gradient Rewrite with Variable

Gradient rewrite algorithrm only follows the data flow, not control edges. So,
it can safely ignore variable write.

However, variable write might mutate the variable value. For example

    # only w is variable
    w = [1, 1]
    a = b * w             # C_1
    w = [2. 2]
    c = a * w             # C_2
    r = reduce(c)
    g = gradient(r, b)
    w = [3, 3]            # C_3

in which we can observe several issues:

1. During the backward pass, the value of `w`, used at `C_1` is changed, we
   need to keep a copy, i.e., snapshot, of varialbe of `w`.
2. We do not need the snapshot of `w` at `C_2` as there is no variable write
   found from `C_2` to gradient backprop.
3. If we use the value of `w` at `C_2` in backward pass, we need to ensure
   the write at `C_3` is after the value use in gradient backprop.

