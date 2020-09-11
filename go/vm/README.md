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
OpLOAD     0  # load 0 (tensor)
OpLOAD     0  # load 0 (tensor)

OpTADD
```

## Example 3 (Linear Regression)

```
# all w_i and x_i are scalar, i.e., elements in a vectors.
# n is a i.i.d noises sampled from standard normal distribution.
y = w_1 * x_1 + w_2 * x_2 + w_3 * x_3 + n
```

```
# Constants
# 0 object.String("w")
# 1 object.String("b")
# 2 object.Shape(<3, 1>)
# 3 object.Integer(456)

OpIOR      2  # read (x,y) for current batch.
OpSTORE    0  # store y and put aside


OpCONST    0  # param key "w"
OPLOADS       # pull "w" from key-value store
OpTMATMUL     # o_1 = matmul(x, w)
OpSTORE    1  # store o_1

OpCONST    1  # param key "b"
OPLOADS       # pull "b" from key-value store
OpLOAD     1
OpTSHAPE      # get shape from o_1
OpTBROAD      # align shape

OpMOVE     1  # move o_1 back
OpTADD        # o_2 = o_1 + b

OpMOVE     0  # move y to stack top
OpTMINUS      # diff_o = y - o_2

OpSTORE    0
OpLOAD     0
OpMOVE     0  # basically dup the diff_o

OpTREDUCE  0  # loss = reduce_sum(diff_o)


# Code
```

