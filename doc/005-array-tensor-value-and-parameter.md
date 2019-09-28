Topic: Array, Tensor, Value, And Parameter.
===========================================

- Draft: 2019-09
- State: Draft

For constant literay, we define an `Array`, like numpy array. It holds the data
with shape. It is immutable given it is a constant.

Tensor should be a concept in `Computation`. It could be `Array`, computation
result, or persistant `Parameter`.


For each Array, user should register it into the Computation so it can be used
with computation later. The ownership of the `Array` will be thereby transfered
to the Computation.

In future, it might be more reasonable to define context to register the Tensor
and Module to define local computation.
