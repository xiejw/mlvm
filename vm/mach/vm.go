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
	if opt != nil {
		opt = opt.Clone() // make a copy
	}

	// ---------------------------------------------------------------------------
	// deduce output dtypes and shapes.
	//
	// In addition, validation will be perfomed in op.OutputTypes
	// ---------------------------------------------------------------------------
	operand_tensor_likes := make([]object.TensorLike, 0, len(operands))
	operand_tensors := make([]*object.Tensor, 0, len(operands))
	for _, opr := range operands {
		operand_tensors = append(operand_tensors, opr.tensor)
		operand_tensor_likes = append(operand_tensor_likes, opr.tensor)
	}

	output_tensor_likes, err := op.OutputTypes(operand_tensor_likes, opt)
	if err != nil {
		return nil, errors.WrapNote(err, "failed to verify op signature during executing op.")
	}

	var outputs []*Handle
	var output_tensors []*object.Tensor
	if len(output_tensor_likes) > 0 {
		outputs = make([]*Handle, 0, len(output_tensor_likes))
		for i := 0; i < len(output_tensor_likes); i++ {
			o, err := vm.NewHandle(output_tensor_likes[i].DType(), output_tensor_likes[i].Shape().Dims)
			if err != nil {
				return nil, errors.WrapNote(err, "failed to allocate output space during executing op.")
			}
			outputs = append(outputs, o)
			output_tensors = append(output_tensors, o.tensor)
		}
	}

	// ---------------------------------------------------------------------------
	// deduce allowing gradients.
	// ---------------------------------------------------------------------------
	flowGrad := vm.shouldFlowGrad(op, operands)

	if flowGrad {
		err := op.AllowGrad(operand_tensor_likes, opt)
		if err != nil {
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
	err = vm.scheduleOp(op, operand_tensors, output_tensors, opt)
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
