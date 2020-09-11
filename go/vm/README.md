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
# @w object.String("w")
# @b object.String("b")
# @lr object.Tensor(<1> [0.001])

# Code

%tmp
%x   1
%y   2
%w   3
%b   4
%o1  5
%do  6
%lo  7
%db  8
%dw  9

# reads input batch

OpIOR      2     # read (x,y) for current batch.
OpSTORE    %y    # store y and put aside
OpSTORE    %x    # store x and put aside

# reads model params

OpCONST    @w    # param key "w"
OPLOADS          # pull "w" from key-value store
OpSTORE    %w

OpCONST    @b     # param key "b"
OPLOADS          # pull "b" from key-value store
OpSTORE    %b

# forward pass

OpLOAD     %x
OpLOAD     %w
OpTMATMUL        # o1 = matmul(x, w)
OpSTORE    %o1   # store o1

OpLOAD     %o1
OpTSHAPE         # get shape from o1
OpTBROAD         # align shape

OpMOVE     %o1   # move o1 back
OpTADD           # o2 = o1 + b

OpMOVE     %y    # move y to stack top
OpTMINUS         # diff_o = y - o2
OpSTORE    %do

OpLOAD     %do
OpMOVE     %do   # basically dup the diff_o

OpTREDUCE  0     # loss = reduce_sum(diff_o)
OpSTORE    %lo

# backprop pass

OpLOAD     %do
OPTREDUCE  0     # grad for w
OpSTORE    %db

# optimizer apply

OpLOAD     %b
OpCONST    @lr
OpLOAD     %db
OpTMUL
OpTMINUS         # b = b - lr * grad_b
OpSTORE    %b

OpLOAD     %w
OpCONST    @lr
OpLOAD     %dw
OpTMUL
OpTMINUS         # w = w - lr * grad_w
OpSTORE    %w


# store back new params
OpCONST    @b
OpMOVE     %b
OpSTORES

OpCONST    @w
OpMOVE     %w
OpSTORES


```

