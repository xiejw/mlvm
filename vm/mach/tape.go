package mach

import (
	"fmt"

	"github.com/xiejw/mlvm/vm/algorithms/autograd"
	"github.com/xiejw/mlvm/vm/ops"
)

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

func (t *Tape) BProp(x *Handle) error {
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
