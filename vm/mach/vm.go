package mach

import (
	"github.com/xiejw/mlvm/vm/algorithms/rngs"
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
	operand_tensors := make([]*object.Tensor, 0, len(operands))
	for _, opr := range operands {
		operand_tensors = append(operand_tensors, opr.tensor)
	}

	output_dtypes, output_shapes, err := op.OutputTypes(operand_tensors, opt)
	if err != nil {
		return nil, errors.WrapNote(err, "failed to verify op signature during executing op.")
	}

	var outputs []*Handle
	if len(output_dtypes) > 0 {
		outputs = make([]*Handle, 0, len(output_dtypes))
		for i := 0; i < len(output_dtypes); i++ {
			o, err := vm.NewHandle(output_dtypes[i], output_shapes[i])
			if err != nil {
				return nil, errors.WrapNote(err, "failed to allocate output space during executing op.")
			}
			outputs = append(outputs, o)
		}
	}

	// ---------------------------------------------------------------------------
	// deduce gradients.
	// ---------------------------------------------------------------------------
	err = vm.validateGradAndRecordOp(op, operands, outputs, opt)
	if err != nil {
		return nil, errors.WrapNote(err, "failed to validate grads and record op during executing op.")
	}

	// ---------------------------------------------------------------------------
	// schedule op.
	err = vm.scheduleOp(op, operands, outputs, opt)
	if err != nil {
		return nil, errors.WrapNote(err, "failed to schedule during executing op.")
	}
	return outputs, nil
}

func (vm *VM) validateGradAndRecordOp(op ops.OpCode, operands []*Handle, outputs []*Handle, opt ops.Option) error {
	flowGrad, err := vm.validateFlowingGradient(op, operands)
	if err != nil {
		return err
	}

	r := &Record{
		Op:       op,
		Operands: operands,
		Outputs:  outputs,
		Option:   opt,
		FLowGrad: flowGrad,
	}

	for _, o := range outputs {
		o.record = r
		o.flowGrad = flowGrad
	}

	vm.tape.Records = append(vm.tape.Records, r)
	return nil
}

func (vm *VM) scheduleOp(op ops.OpCode, operands []*Handle, outputs []*Handle, opt ops.Option) error {
	// TODO: use async

	switch op {
	case ops.OP_RNG:
		// verified by validateSignature
		value := operands[0].tensor.Data.([]float32)
		rng_opt := opt.(*ops.RngOption)

		switch rng_opt.DistType {
		case ops.RngDistStdNorm:
			rngs.StdNorm(rng_opt.Rng, value)
			return nil
		case ops.RngDistTruncStdNorm:
			rngs.TruncStdNorm(rng_opt.Rng, value)
			return nil
		default:
			return errors.New("unknown distribution type: %v", rng_opt.DistType)
		}
	default:
		return errors.New("unsupported op (%v) for scheduling op", op)
	}
}

func (vm *VM) validateFlowingGradient(op ops.OpCode, operands []*Handle) (bool, error) {
	flowGrad := false
	for _, opr := range operands {
		if opr.flowGrad || opr.requireGrad {
			flowGrad = true
			break
		}
	}

	if !flowGrad {
		return false, nil
	}

	// TODO must F32 for flowGrad

	switch op {
	case ops.OP_RNG:
		err := errors.New("op (%v) cannot flow grad.", op)

		// emits more info
		for i, opr := range operands {
			if opr.flowGrad {
				err.EmitNote("the %v-th operand needs to flow grad.", i)
				break
			}
			if opr.requireGrad {
				err.EmitNote("the %v-th operand requires grad.", i)
				break
			}
		}
		return false, err
	default:
	}
	return flowGrad, nil
}
