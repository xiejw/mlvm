// Tensor implemention in vm.
package tensorarray

import (
	"github.com/xiejw/mlvm/go/object"
)

type TensorArray struct {
	Dims     []int
	Rank     int
	Size     int
	RealSize int
	Value    []float32
}

// Creates TensorArray from raw components.
func FromRaw(dims []int, value []float32) *TensorArray {
	rank := len(dims)

	var size = 1
	for _, dim := range dims {
		size *= dim
	}

	return &TensorArray{
		Dims:     dims,
		Rank:     rank,
		Size:     size,
		RealSize: len(value),
		Value:    value,
	}
}

// Helper Method to create TensorArray from Tensor.
func FromTensor(t *object.Tensor) *TensorArray {
	return FromRaw(t.Shape.Dims, t.Array.Value)
}

func (ta *TensorArray) ToTensor() *object.Tensor {
	if ta.IsCompressed() {
		panic("Converting compressed TensorArray to Tensor is not impl'ed.")
	}

	return object.NewTensor(ta.Dims, ta.Value)
}

func (ta *TensorArray) IsCompressed() bool {
	return ta.Size != ta.RealSize
}

// Conform object.Object
func (ta *TensorArray) String() string {
	return "TensorArray"
}

func (ta *TensorArray) Type() object.ObjectType {
	return object.TensorType
}
