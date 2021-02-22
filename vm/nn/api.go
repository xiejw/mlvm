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
	operands := []*mach.Handle{lhs, rhs}
	vm := assertSameVM(operands)
	handle, err := vm.ExecOp(ops.OP_ADD, operands, nil)
	if err != nil {
		panic(err)
	}

	return handle
}

func Mul(lhs, rhs *mach.Handle) *mach.Handle {
	operands := []*mach.Handle{lhs, rhs}
	vm := assertSameVM(operands)
	handle, err := vm.ExecOp(ops.OP_MUL, operands, nil)
	if err != nil {
		panic(err)
	}

	return handle
}

func Sum(x *mach.Handle) *mach.Handle {
	operands := []*mach.Handle{x}
	vm := assertSameVM(operands)
	handle, err := vm.ExecOp(ops.OP_SUM, operands, &ops.SumOption{x.Shape().Dims})
	if err != nil {
		panic(err)
	}

	return handle
}

func Backward(x *mach.Handle) {
}

// -----------------------------------------------------------------------------
// helper methods.
// -----------------------------------------------------------------------------
func assertSameVM(operands []*mach.Handle) *mach.VM {
	switch len(operands) {
	case 0:
		panic("expected non-emtpy operands slice passed in.")
	case 1:
		return operands[0].VM()
	default:
		vm := operands[0].VM()
		return vm
	}
}
