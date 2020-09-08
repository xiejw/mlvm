// Tensor implemention in vm.
package tensorarray

import (
	"github.com/xiejw/mlvm/go/object"
)

type TensorArray struct {
	Dims       []uint
	Strides    []uint
	Rank       uint
	Value      []float32
	Compressed bool
}

func FromTensor(t *object.Tensor) *TensorArray {
	dims := t.Shape.Dims
	rank := t.Shape.Rank
	strides := make([]uint, rank)

	var stride uint = 1
	for i := int(rank - 1); i >= 0; i-- {
		strides[i] = stride
		stride *= dims[i]
	}

	return &TensorArray{
		Dims:       dims,
		Strides:    strides,
		Rank:       rank,
		Value:      t.ArrayValue(),
		Compressed: false,
	}
}

func (ta *TensorArray) ToTensor() *object.Tensor {
	if ta.Compressed {
		panic("Converting compressed TensorArray to Tensor is not impl'ed.")
	}

	return object.NewTensor(ta.Dims, ta.Value)
}

// Conform object.Object
func (ta *TensorArray) String() string {
	return "TensorArray"
}

func (ta *TensorArray) Type() object.ObjectType {
	return object.TensorType
}
