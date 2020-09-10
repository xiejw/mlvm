package kernel

import (
	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/vm/tensorarray"
)

type BinaryOpType int

const (
	BinaryAdd BinaryOpType = iota
	BinaryMinus
	BinaryMul
)

// Algorithrm for the binary Ops.
//
// Input Requirments: The shape must match exactly. For program writer, use OpTBROAD if needed.
//
// 1. If both the size and real_size are same, then performn buffer operation directly.
// 2. otherwise,, then using a recursive loop to form binary op in each dim.
func BinaryOp(o1, o2 *tensorarray.TensorArray, op_type BinaryOpType) (
	*tensorarray.TensorArray, *errors.DError) {

	if !areUIntSliceEq(o1.Dims, o2.Dims) {
		return nil, errors.New("dims in shape mismatch. this is not allowed.")
	}

	var buf []float32

	if o1.RealSize == o2.RealSize {
		// Simple case. Perform operations directly.
		size := o1.Size
		buf1 := o1.Value
		buf2 := o2.Value

		buf = make([]float32, size)

		switch op_type {
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
			err := errors.New("unsupported binary op type %v", op_type)
			return nil, err
		}
	} else if o1.RealSize == 1 {
		real_size_2 := o2.RealSize
		buf2 := o2.Value

		buf = make([]float32, real_size_2)
		v := o1.Value[0]

		switch op_type {
		case BinaryAdd:
			for i := 0; i < real_size_2; i++ {
				buf[i] = v + buf2[i]
			}
		case BinaryMinus:
			for i := 0; i < real_size_2; i++ {
				buf[i] = v - buf2[i]
			}
		case BinaryMul:
			for i := 0; i < real_size_2; i++ {
				buf[i] = v * buf2[i]
			}
		default:
			err := errors.New("unsupported binary op type %v", op_type)
			return nil, err
		}
	} else if o2.RealSize == 1 {
		real_size_1 := o1.RealSize
		buf1 := o1.Value

		buf = make([]float32, real_size_1)
		v := o2.Value[0]

		switch op_type {
		case BinaryAdd:
			for i := 0; i < real_size_1; i++ {
				buf[i] = buf1[i] + v
			}
		case BinaryMinus:
			for i := 0; i < real_size_1; i++ {
				buf[i] = buf1[i] - v
			}
		case BinaryMul:
			for i := 0; i < real_size_1; i++ {
				buf[i] = buf1[i] * v
			}
		default:
			err := errors.New("unsupported binary op type %v", op_type)
			return nil, err
		}

	} else {

		// real_size_1 := o1.RealSize
		// real_size_2 := o2.RealSize

		// max_size := real_size_1
		// min_size := real_size_2

		// if max_size < min_size {
		// 	max_size = real_size_2
		// 	min_size = real_size_1
		// }

		// if max_size%min_size != 0 {
		// 	return nil, errors.New(
		// 		"real size of operands are invalid. they should be multiple of each other: got: %v and %v",
		// 		real_size_1, real_size_2)
		// }

		// for i := 0; i < max_size; i++ {
		// }

	}

	return tensorarray.FromRaw(o1.Dims, buf), nil
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
