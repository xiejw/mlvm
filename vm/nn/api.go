package nn

import (
	"github.com/xiejw/mlvm/vm/algorithms/rngs"
	"github.com/xiejw/mlvm/vm/mach"
	"github.com/xiejw/mlvm/vm/object"
)

func RngStdNormal(vm *mach.VM, rng rngs.Rng, dtype object.DType, dims []int) *mach.Handle {
	return nil
}

func Add(lhs, rhs *mach.Handle) *mach.Handle {
	return nil
}

func Mul(lhs, rhs *mach.Handle) *mach.Handle {
	return nil
}

func Sum(x *mach.Handle) *mach.Handle {
	return nil
}

func Backward(x *mach.Handle) {
}
