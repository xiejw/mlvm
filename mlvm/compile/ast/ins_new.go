package ast

import (
	"github.com/xiejw/mlvm/mlvm/internal/naming"
)

func newInstruction(name string, op *Op, operands ...*Tensor) *Instruction {

	if err := naming.ValidateInstructionName(name); err != nil {
		panic(err)
	}

	ins := &Instruction{
		name:     name,
		op:       op,
		operands: operands,
	}

	// TODO: Remove the hard code.
	index := 0
	resultName := naming.CanonicalResultName(name, index)

	result := &Result{
		name:  resultName,
		shape: operands[0].Shape(),
		ins:   ins,
		index: index,
	}

	ins.results = []*Tensor{newResultTensor(result)}
	return ins
}
