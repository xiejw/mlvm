package ast

func newInstruction(name string, op *Op, operands ...*Tensor) *Instruction {
	ins := &Instruction{
		name:     name,
		op:       op,
		operands: operands,
	}

	// TODO: Remove the hard code.
	result := &Result{
		name:  "%o_0",
		shape: operands[0].Shape(),
		ins:   ins,
		index: 0,
	}

	ins.results = []*Tensor{newResultTensor(result)}
	return ins
}
