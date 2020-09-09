package kernel

import (
	"github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/vm/tensorarray"
)

// Algorithrm for the binary Ops.
//
// Input Requirments: The shape must match exactly. For program writer, use OpTBROAD if needed.
//
// 1. If the strides are same, then performn buffer adding directly.
// 2. If the strides are not same, then using a recursive loop to form add in each dim.
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

func isShapeEqual() bool {
	return true
}
