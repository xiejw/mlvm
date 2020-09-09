package kernel

import (
	"github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/vm/tensorarray"
)

// Algorithrm for the binary Ops.
//
// Input Requirments:
// 1. The
func TensorAdd(o1, o2 *tensorarray.TensorArray) (*tensorarray.TensorArray, error) {
	operand1 := o1.ToTensor()
	operand2 := o2.ToTensor()
	shape := operand1.Shape
	size := shape.Size()

	buf := make([]float32, size)

	buf1 := operand1.ArrayValue()
	buf2 := operand2.ArrayValue()

	var i uint64
	for i = 0; i < size; i++ {
		buf[i] = buf1[i] + buf2[i]
	}

	return tensorarray.FromTensor(&object.Tensor{shape, &object.Array{buf}}), nil
}
