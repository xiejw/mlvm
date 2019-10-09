package ast

import (
	"github.com/xiejw/mlvm/mlvm/internal/naming"
)

// Creates a new Instruction in Module.
func (m *Module) NewInstruction(op *Op, operands ...*Tensor) *Instruction {
	baseName := op.BaseName()
	var name string
	for {
		m.opNameIndex += 1
		name = naming.DefaultInstructionName(baseName, m.opNameIndex)
		if m.nameStore[name] == nil {
			break
		}
	}
	return m.NewInstructionWithName(name, op, operands...)
}

func (m *Module) NewInstructionWithName(name string, op *Op, operands ...*Tensor) *Instruction {
	m.mustNotFreezed()

	ins := newInstruction(name, op, operands...)

	m.registerName(name, ins, true /* registerOnce */)
	m.instructions = append(m.instructions, ins)
	return ins
}
