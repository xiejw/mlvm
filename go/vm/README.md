### Example 1

Adds two `Tensors` created from constants.

```
# Constants
# 0 object.Shape(<2, 3>)
# 1 object.Array([1.0, 2.0, 3.0, 4.0, 5.0, 6.0])

# Code
OpCONST    0
OpCONST    1
OpT            # create 1st operand.

OpCONST    0
OpCONST    1
OpT            # create 2nd operand.

OpTADD
```

### Example 2


```
# Constants
# 0 object.Shape(<2, 3>)
# 1 object.Integer(456)

# Code
OpCONST    0  # shape
OpCONST    1  # seed int
OpRNG

OpRNGT     0  # tensor with noraml dist
OpSTORE    0  # store to 0
OpCONST    0  # shape
OpLOAD     0  # load 0 (array)
OpT

OpSTORE    0  # store to 0
OpLOAD     0  # load 0 (tensor)
OpLOAD     0  # load 0 (tensor)

OpAdd
```
