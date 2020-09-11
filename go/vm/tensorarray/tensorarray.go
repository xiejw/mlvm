// Tensor implemention in vm.
package tensorarray

import (
	"github.com/xiejw/mlvm/go/object"
)

// TensorArray supports one simple form of broadcasting, on top of the normal Tensor representation.
//
// For any shape [a1, a2, a3] and a1 * a2 * a3 elements for a tensor, it is called non-compressed
// tensor array, as all elements have been filled in value buffer.
//
// The only valid broadcasting case is [b1, b2, b3, a1, a2, a3], i.e., major dimension extension.
// For example, all of the following cases are supported
//
//   - yes [2, 3] -> [3, 2, 3]
//   - yes [2, 3] -> [4, 3, 2, 3]
//
//   - no [2, 3] -> [4, 3, 2, 1]   the final dim can only be 3, but got 1.
//   - no [2, 3] -> [3, 4, 3]      the second from last dim can only be 2, but got 4.
//
// As a special but super useful use case, shape [1] is allowed to be broadcasted to all other
// shapes, i.e.,
//
//   - yes [1] -> [3, 2, 1]
//   - yes [1] -> [4, 3, 2, 3]
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
