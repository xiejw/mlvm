package mach

import (
	"fmt"

	"github.com/xiejw/mlvm/vm/algorithms/autograd"
	"github.com/xiejw/mlvm/vm/base/errors"
	"github.com/xiejw/mlvm/vm/ops"
)

// most are expected to be immutable. few exceptions are
//
// - handle passed to non-grad-flow op which later becomes required grad.
//   handle -> rng -> require_grad.
type Record struct {
	Op       ops.OpCode
	Operands []*Handle
	Outputs  []*Handle
	Option   ops.Option
}

type Tape struct {
	Records []*Record
	GradDAG []*Record
}

func (t *Tape) RecordOpAndGradDAG(
	op ops.OpCode, operands []*Handle, outputs []*Handle, opt ops.Option, flowGrad bool,
) error {
	r := &Record{
		Op:       op,
		Operands: operands,
		Outputs:  outputs,
		Option:   opt,
	}

	for _, o := range outputs {
		o.record = r
		o.flowGrad = flowGrad
	}

	t.Records = append(t.Records, r)
	if flowGrad {
		t.GradDAG = append(t.GradDAG, r)
	}
	return nil
}

// algorithrm
// - prepare the ones tensor matching the shape of passed in x.
// - prepare a map, called handle_grad, from handle to current grad. start with the x -> ones
//   created in last step.
// - walk through dag in topological reverse order. iterate on each Op.
//   - look up the gradients for outputs handles, via handle_grad. error if cannot find.
//   - call autograd to get the gradients.
//   - for each operand
//     - discard gradient if neither require grad nor flow grad.
//     - if flow grad, add the current grad on top of the old one in handle_grad and save new result
//       to handle_grad
//     - if require_grad, add the current grad on top of the handle'grad handle.
func (t *Tape) BProp(x *Handle) error {
	if !x.flowGrad {
		return errors.New(
			"Handle (%v) has not grad flowing through it; so cannot be used for Backward", x)
	}

	fmt.Printf("tape: \n")
	for _, r := range t.Records {
		fmt.Printf("  %+v\n", r)
	}
	fmt.Printf("dag: \n")
	for _, r := range t.GradDAG {
		fmt.Printf("  %+v\n", r)
	}

	for _, r := range t.GradDAG {
		_, err := autograd.Grad(
			r.Op,
			r.Option,
			handlesToTensorLikes(r.Operands),
			handlesToTensorLikes(r.Outputs),
			nil)
		if err != nil {
			return err
		}

	}
	return nil
}
