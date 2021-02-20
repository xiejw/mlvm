package mach

import (
	// "github.com/xiejw/mlvm/vm/base/errors"
	"github.com/xiejw/mlvm/vm/object"
	"github.com/xiejw/mlvm/vm/ops"
)

type VM struct {
	tape    Tape
	handles []*Handle
}

func (vm *VM) ExecOp(op ops.OpCode, operands []*Handle, opt ops.Option) (*Handle, error) {
	//	func FillDist(rng Rng, distType DistType, value []float32) {
	//		switch distType {
	//		case DistStdNorm:
	//			StdNorm(rng, value)
	//			return
	//		case DistTruncStdNorm:
	//			TruncStdNorm(rng, value)
	//			return
	//		default:
	//			panic(fmt.Sprintf("unknown distribution type: %v", distType))
	//		}
	//	}
	return nil, nil
}

func (vm *VM) NewHandle(dtype object.DType, dims []int) (*Handle, error) {

	handle := &Handle{
		id:          len(vm.handles),
		tensor:      object.NewTensor(dtype, dims),
		vm:          vm,
		requireGrad: false,
		gradHandle:  nil,
	}

	vm.handles = append(vm.handles, handle)
	return handle, nil
}
