### Example 1

User code
```
@x = 2;
@y = 3;

let shape = shape_create([@x, @y]);
let value = array_create([1.0, 2.0, 3.0, 4.0, 5.0, 6.0]);
let t1 = tensor_create(shape=shape, value=value);
let t2 = tensor_create(shape=shape, value=value);
let t3 = t1 + t2;
```

VM Compiled Code

```
# Constants
# 0 object.Shape(@x:2, @y:3)
# 1 object.Array([1.0, 2.0, 3.0, 4.0, 5.0, 6.0])

# Code
OpConstant 0
OpConstant 1
OpTensor
OpConstant 0
OpConstant 1
OpTensor
OpAdd
```

### Example 2

User code
```
@x = 2
@y = 3

let shape = shape_create([@x, @y]);
let prng = prng_create(seed: 456);
let value = prng_normal(source: prng, shape: shape);
let t1 = tensor_create(shape=shape, value=value)
let t2 = tensor_create(shape=shape, value=value)
let t3 = t1 + t2
```

VM Compiled Code

```
# Constants
# 0 object.Shape(@x:2, @y:3)
# 1 object.Integer(456)

# Code
OpConstant 0  # shape
OpConstant 1  # seed int
OpPrngNew
OpPrngNorm    # array
OpStoreG   0  # store to 0
OpConstant 0  # shape
OpLoadG    0  # load 0 (array)
OpTensor
OpConstant 0  # shape
OpLoadG    0  # load 0 (array)
OpTensor
OpAdd
```
