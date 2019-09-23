Topic: Variable Write Hoist
===========================

- Draft: 2019-08
- State: Abandoned

Another Idea
------------

Another idea, which might be more user friendly, is to hoist all variable writes
to the end of computation.

The semantics are defined as follows

1. The read after a write observes the new value, during the computation.
2. The variable value is guaranteed to be updated by the end of the computation.
   This means we are allowed to defer variable updates.

### Example Study

Consider the following computation written by users.

    # only w is variable
    w = [1, 1]
    a = b * w
    w = [2, 2]
    c = a * w
    r = reduce(c)
    g = gradient(r, b)

We could safely rewrite this as


    # only w is variable
    c_1 = [1, 1]
    a = b * c_1
    c_2 = [2, 2]
    c = a * c_2
    r = reduce(c)
    g = gradient(r, b)
    w = c_2


Something More
--------------

### Gradient Rewrite

With this rewrite rule, the gradient lowering is safe and much easier. We could
reply on optimization pass later to save memory due to extra copy.

### Dead Nodes

Consider the following computation written by users.

    # only w is variable
    a = c_1 * w
    w = [2, 2]
    c = c_2 * w
    r = reduce(c)
    g = gradient(r, b)

Based computation symantics, the assignment of `a` is a dead node. If we perform
the variable write hoist rewrite, the result graph looks like

    a = c_1 * read(w)
    c_3 = [2, 2]
    c = c_2 * c_3
    r = reduce(c)
    g = gradient(r, b)
    w = c_3

This is clear enough that we could remove that branch easily.

### Illegal Graph

At the moment, it is not clear how to solve this issue.
