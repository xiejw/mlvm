package nn

import (
	"github.com/xiejw/mlvm/vm/algorithms/rngs"
	"github.com/xiejw/mlvm/vm/mach"
	"github.com/xiejw/mlvm/vm/object"
	"github.com/xiejw/mlvm/vm/ops"
)

func RngStdNorm(vm *mach.VM, rng rngs.Rng, dtype object.DType, dims []int) *mach.Handle {
	handle, err := vm.NewHandle(dtype, dims)
	if err != nil {
		panic(err)
	}

	_, err = vm.ExecOp(ops.OP_RNG, []*mach.Handle{handle}, &ops.RngOption{rng, ops.RngDistStdNorm})
	if err != nil {
		panic(err)
	}

	return handle
}

func Zeros(vm *mach.VM, dtype object.DType, dims []int) *mach.Handle {
	return nil
}

func Ones(vm *mach.VM, dtype object.DType, dims []int) *mach.Handle {
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
