package mach

import (
	"github.com/xiejw/mlvm/vm/base/errors"
	"github.com/xiejw/mlvm/vm/object"
	"github.com/xiejw/mlvm/vm/ops"
)

type VM struct {
	tape    Tape
	handles []*Handle
}

func (vm *VM) ExecOp(op ops.OpCode, operands []*Handle, opt ops.Option) (*Handle, error) {
	outputs, err := vm.execOp(op, operands, opt)
	if err != nil {
		return nil, err
	}
	switch len(outputs) {
	case 0:
		return nil, nil
	case 1:
		return outputs[0], nil
	default:
		return nil, errors.New("expect returning no more than one output for op (%v), but got: %v", op, len(outputs))
	}
}

func (vm *VM) execOp(op ops.OpCode, operands []*Handle, opt ops.Option) ([]*Handle, error) {
	opt = opt.Clone() // make a copy
	outputs, err := vm.allocateOutputs(op, operands, opt)
	if err != nil {
		return nil, errors.WrapNote(err, "failed to allocate outputs during executing op.")
	}

	err = vm.recordOp(op, operands, outputs, opt)
	if err != nil {
		return nil, errors.WrapNote(err, "failed to record during executing op.")
	}

	err = vm.scheduleOp(op, operands, outputs, opt)
	if err != nil {
		return nil, errors.WrapNote(err, "failed to schedule during executing op.")
	}
	return outputs, nil
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

func (vm *VM) allocateOutputs(op ops.OpCode, operands []*Handle, opt ops.Option) ([]*Handle, error) {
	switch op {
	case ops.OP_RNG:
		return nil, nil
	default:
		return nil, errors.New("unsupported op (%v) for allocating outputs", op)
	}
}

func (vm *VM) recordOp(op ops.OpCode, operands []*Handle, outputs []*Handle, opt ops.Option) error {
	// TODO flow grad.
	vm.tape.Records = append(vm.tape.Records, &Record{
		Op:       op,
		Operands: operands,
		Outputs:  outputs,
		Option:   opt,
		FLowGrad: false,
	})
	return nil
}

func (vm *VM) scheduleOp(op ops.OpCode, operands []*Handle, outputs []*Handle, opt ops.Option) error {
	switch op {
	//	case ops.OP_RNG:
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
	default:
		return errors.New("unsupported op (%v) for scheduling op", op)
	}
}
