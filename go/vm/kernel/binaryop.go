package kernel

import (
	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/vm/tensorarray"
)

type KernelType int

const (
	BinaryAdd KernelType = iota
	BinaryMinus
	BinaryMul
)

// Algorithrm for the binary Ops.
//
// Input Requirments: The shape must match exactly. For program writer, use OpTBROAD if needed.
//
// 1. If the strides are same, then performn buffer operation directly.
// 2. If the strides are not same, then using a recursive loop to form binary op in each dim.
func BinaryOp(o1, o2 *tensorarray.TensorArray, kern_type KernelType) (
	*tensorarray.TensorArray, *errors.DError) {

	if !areUIntSliceEq(o1.Dims, o2.Dims) {
		return nil, errors.New("dims mismatch.")
	}

	if !areUIntSliceEq(o1.Strides, o2.Strides) {
		return nil, errors.New("strides mismatch.")
	}

	operand1 := o1.ToTensor()
	operand2 := o2.ToTensor()
	shape := operand1.Shape
	size := shape.Size()

	buf := make([]float32, size)

	buf1 := operand1.ArrayValue()
	buf2 := operand2.ArrayValue()

	switch kern_type {
	case BinaryAdd:
		for i := 0; i < size; i++ {
			buf[i] = buf1[i] + buf2[i]
		}
	case BinaryMinus:
		for i := 0; i < size; i++ {
			buf[i] = buf1[i] - buf2[i]
		}
	case BinaryMul:
		for i := 0; i < size; i++ {
			buf[i] = buf1[i] * buf2[i]
		}
	default:
		err := errors.New("unsupported Kernel type %v", kern_type)
		return nil, err
	}

	return tensorarray.FromRaw(shape.Dims, buf), nil
}

func areUIntSliceEq(d1, d2 []int) bool {
	l1 := len(d1)
	if l1 != len(d2) {
		return false
	}

	for i := 0; i < l1; i++ {
		if d1[i] != d2[i] {
			return false
		}
	}
	return true
}
