package ast

import ()

// Creates a new Instruction in Module.
func (m *Module) NewInstruction(name string, op *Op, operands ...*Tensor) *Instruction {
	ins := &Instruction{
		name:     name,
		op:       op,
		operands: operands,
	}
	m.registerName(name, ins, true /* registerOnce */)
	return ins
}
