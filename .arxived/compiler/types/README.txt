The Type System

MLVM has a quite simple type system.

The type for basic builtins is



## Strong Shape in Type

For this approach, we only accept Tensor<3x3> as a valid shape, not Tensor. It
gives extremely good debugging ability but it might not be easy.

The way to make it work is likely

1. Each function is like a c++ template. It will be initialized only when the
   input type is known.
2. During invokcation, we initialize the function.

Think

```code
(fn [f g mode::Bool]
    (if mode f g)
)
```

Here, there is no way to do shape inference for both f and g as only one
function will be invoked. Then the return type of the if is not guaranteed to
be the same for all possible mode. So, break strong type rule.

This approach also makes the whole codbase written in a way without any type as
not feasiable.

## Tensor in Function

Allows Tensor in Func and enforce TensorShape in all lead nodes (literal, readIO, etc).

