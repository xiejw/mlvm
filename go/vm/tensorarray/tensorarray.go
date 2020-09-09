// Tensor implemention in vm.
package tensorarray

import (
	"github.com/xiejw/mlvm/go/object"
)

type TensorArray struct {
	Dims       []int
	Strides    []int
	Rank       int
	Value      []float32
	Compressed bool
}

func (ta *TensorArray) Size() int {
	var size = 1
	for _, dim := range ta.Dims {
		size *= dim
	}
	return size
}

func (ta *TensorArray) RealSize() int {
	return len(ta.Value)
}

func FromTensor(t *object.Tensor) *TensorArray {
	dims := t.Shape.Dims
	rank := t.Shape.Rank
	strides := make([]int, rank)

	var stride = 1
	for i := rank - 1; i >= 0; i-- {
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
