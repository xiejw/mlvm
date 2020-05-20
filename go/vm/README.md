Use code
```
@x = 2
@y = 3

let shape: Shape = Shape([@x, @y]);
let value: Array = [1.0, 2.0, 3.0, 4.0, 5.0, 6.0];
let t1: Tensor = Tensor(shape=shape, value=value)
let t2: Tensor = Tensor(shape=shape, value=value)
let t3: Tensor = t1 + t2
```

VM Compiled Code

```
# Constant
# 0 Shape(@x:2, @y:3)
# 1 Array([1.0, 2.0, 3.0, 4.0, 5.0, 6.0])

# Code
OpConstant 0
OpConstant 1
OpTensor
OpConstant 0
OpConstant 1
OpTensor
OpAdd
```
