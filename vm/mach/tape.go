package mach

import (
	"fmt"

	"github.com/xiejw/mlvm/vm/ops"
)

type Record struct {
	Op       ops.OpCode
	Operands []*Handle
	Outputs  []*Handle
	Option   ops.Option
	FLowGrad bool
}

type Tape struct {
	Records []*Record
}

func (t *Tape) RecordOpAndGradDAG(
	op ops.OpCode, operands []*Handle, outputs []*Handle, opt ops.Option, flowGrad bool,
) error {
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

	t.Records = append(t.Records, r)
	return nil
}

func (t *Tape) BProp(x *Handle) error {
	fmt.Printf("tape: \n")
	for _, r := range t.Records {
		fmt.Printf("  %+v\n", r)
	}
	return nil
}
