// Tensor implemention in vm.
package tensorarray

import (
	"github.com/xiejw/mlvm/go/object"
)

type TensorArray struct {
	Dims   []uint
	Strides []uint
	Rank   uint
	Value  []float32
}

func FromTensor(t *object.Tensor) *TensorArray {
	dims:= t.Shape.Dims
	rank := t.Shape.Rank
	strides := make([]uint,rank)

	var stride uint = 1
	for i := int(rank -1) ; i >= 0; i--{
		strides[i] = stride
		stride *= dims[i]
	}

	return &TensorArray{
		Dims: dims,
		Strides: strides,
	  Rank: rank,
		Value: t.ArrayValue(),
	}
}
