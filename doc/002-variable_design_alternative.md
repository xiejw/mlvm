Topic: Another Way to Design the variable
==========================================

- Draft: 2019-08
- State: Implemented

Idea
----

We can view all variables as pure computation (module) inputs. And have
user-facing APIs to write results to write back.

    # Literals
    a := newTensor("a", ...)
    b := newTensor("b", ...)

    # Use default container.
    w := variable.NewVariableFromLiteral("w", ...)

    m := module.NewModule(...)

    c := m.Mul("addw", a, w)
    c_1 := m.Add("addb", c, b)
    d := m.Reshape("reshape", c_1, []tensor.Dimension{2, 2})
    loss := m.Reduce("loss", d, module.ReduceSum)
    g := m.Gradient("gradient", loss, []tensor.Tensor{d, c_1, w})

    gradient := g[2]
    new_w := m.Minus("minus", w, gradient)

    m.SetOutputsWithVariableWrites(
       /* outputs=*/ []tensor.Tensor{c, d, loss, g[0], g[1], g[2]},
       []*VariableWrite {
          {Variable: w, NewValue: new_w},
       })

Given the complexity of optimization and dependency control with variable
read/write interleaving in the computation, it is much easier to get define the
computation.

The Illness of Graph with Variable Write
----------------------------------------

    # only w is variable
    w = [1, 1]
    a = b * w
    w = [2, 2]
    c = a * w
    r = reduce(c)
    g = gradient(r, w)

If we allow variable write in the middle of computation, we would face the ill
graph like above, which does not have meaningful gradient definition.
