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

// -----------------------------------------------------------------------------
// public apis.
// -----------------------------------------------------------------------------

func (vm *VM) NewHandle(dtype object.DType, dims []int) (*Handle, error) {
	handle := &Handle{
		id:          len(vm.handles),
		tensor:      object.NewTensor(dtype, dims),
		vm:          vm,
		requireGrad: false,
		flowGrad:    false,
		gradHandle:  nil,
		record:      nil,
	}

	vm.handles = append(vm.handles, handle)
	return handle, nil
}

func (vm *VM) ExecOp(op ops.OpCode, operands []*Handle, opt ops.Option) (*Handle, error) {
	if opt != nil {
		opt = opt.Clone() // make a copy
	}

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

func (vm *VM) Backward(x *Handle) error {
	return vm.tape.BProp(x)
}

func (vm *VM) WaitBarrier() {
}

// -----------------------------------------------------------------------------
// helper methods.
// -----------------------------------------------------------------------------

func (vm *VM) execOp(op ops.OpCode, operands []*Handle, opt ops.Option) ([]*Handle, error) {

	// ---------------------------------------------------------------------------
	// deduce outputs signatures: dtypes and shapes.
	//
	// In addition, validation will be perfomed in op.InferOutputs
	// ---------------------------------------------------------------------------
	operand_sigs := HandlesToTensorLikes(operands)
	output_sigs, err := op.InferOutputs(operand_sigs, opt)
	if err != nil {
		return nil, errors.WrapNote(err, "failed to verify op signature during executing op.")
	}

	outputs, err := vm.allocateHandlesForOutputs(output_sigs)
	if err != nil {
		return nil, errors.WrapNote(err, "failed to allocate output space during executing op.")
	}

	// ---------------------------------------------------------------------------
	// deduce allowing gradients.
	// ---------------------------------------------------------------------------
	flowGrad := vm.shouldFlowGrad(op, operands)
	if flowGrad {
		if err := op.AllowGrad(operand_sigs, opt); err != nil {
			return nil, errors.WrapNote(err, "failed to flow grad during executing op.")
		}
	}

	// ---------------------------------------------------------------------------
	// recode op and gradent dag.
	// ---------------------------------------------------------------------------
	err = vm.tape.RecordOpAndGradDAG(op, operands, outputs, opt, flowGrad)
	if err != nil {
		return nil, errors.WrapNote(err, "failed to record op and grad DAG during executing op.")
	}

	// ---------------------------------------------------------------------------
	// schedule op.
	err = vm.scheduleOp(op, HandlesToTensors(operands), HandlesToTensors(outputs), opt)
	if err != nil {
		return nil, errors.WrapNote(err, "failed to schedule during executing op.")
	}
	return outputs, nil
}

func (vm *VM) shouldFlowGrad(op ops.OpCode, operands []*Handle) bool {
	for _, opr := range operands {
		if opr.flowGrad || opr.requireGrad {
			return true
		}
	}
	return false
}

func (vm *VM) scheduleOp(op ops.OpCode, operands []*object.Tensor, outputs []*object.Tensor, opt ops.Option) error {
	// TODO: use async
	return op.Exec(operands, outputs, opt)
}

func (vm *VM) allocateHandlesForOutputs(output_sigs []object.TensorLike) ([]*Handle, error) {
	size := len(output_sigs)
	if size == 0 {
		return nil, nil
	}

	outputs := make([]*Handle, 0, size)
	for i := 0; i < size; i++ {
		o, err := vm.NewHandle(output_sigs[i].DType(), output_sigs[i].Shape().Dims)
		if err != nil {
			return nil, errors.WrapNote(err, "failed to create handle for %v-th output.", i)
		}
		outputs = append(outputs, o)
	}
	return outputs, nil
}

func HandlesToTensors(hs []*Handle) []*object.Tensor {
	size := len(hs)
	ts := make([]*object.Tensor, 0, size)
	for _, h := range hs {
		ts = append(ts, h.tensor)
	}
	return ts
}

func HandlesToTensorLikes(hs []*Handle) []object.TensorLike {
	size := len(hs)
	ts := make([]object.TensorLike, 0, size)
	for _, h := range hs {
		ts = append(ts, h.tensor)
	}
	return ts
}
