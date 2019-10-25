package ir

import (
	"github.com/xiejw/mlvm/mlvm/internal/naming"
)

// Creates a new Instruction in Module.
//
// This is convenient to use given it will try to auto generate the Instruction name and it will
// panic for any error.
func (m *Module) NewInstructionOrDie(op *Op, operands ...*Tensor) *Instruction {
	baseName := op.BaseName()
	var name string
	for {
		m.opNameIndex += 1
		name = naming.DefaultInstructionName(baseName, m.opNameIndex)
		if m.nameStore[name] == nil {
			break
		}
	}
	ins, err := m.NewInstructionWithName(name, op, operands...)
	if err != nil {
		panic(err)
	}
	return ins
}

func (m *Module) NewInstructionWithName(
	name string, op *Op, operands ...*Tensor,
) (*Instruction, error) {

	err := m.mustNotFreezed()
	if err != nil {
		return nil, err
	}

	ins, err := newInstruction(name, op, operands...)
	if err != nil {
		return nil, err
	}

	err = m.registerName(name, ins, true /* registerOnce */)
	if err != nil {
		return nil, err
	}

	m.instructions = append(m.instructions, ins)
	return ins, nil
}
