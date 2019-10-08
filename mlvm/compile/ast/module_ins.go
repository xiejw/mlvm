package ast

import (
	"fmt"
)

// Creates a new Instruction in Module.
func (m *Module) NewInstruction(op *Op, operands ...*Tensor) *Instruction {
	baseName := op.BaseName()
	var name string
	for {
		m.opNameIndex += 1
		name = fmt.Sprintf("%v_%03v", baseName, m.opNameIndex)
		if m.nameStore[name] == nil {
			break
		}
	}
	return m.NewInstructionWithName(name, op, operands...)
}

func (m *Module) NewInstructionWithName(name string, op *Op, operands ...*Tensor) *Instruction {
	m.mustNotFreezed()

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

	m.registerName(name, ins, true /* registerOnce */)
	m.instructions = append(m.instructions, ins)
	return ins
}
