package kernel

import (
	"github.com/xiejw/mlvm/go/object"
)

func TensorAdd(operand1, operand2 *object.Tensor) (*object.Tensor, error) {
	shape := operand1.Shape
	size := shape.Size()

	buf := make([]float32, size)

	buf1 := operand1.ArrayValue()
	buf2 := operand2.ArrayValue()

	var i uint64
	for i = 0; i < size; i++ {
		buf[i] = buf1[i] + buf2[i]
	}

	return &object.Tensor{shape, &object.Array{buf}}, nil
}
