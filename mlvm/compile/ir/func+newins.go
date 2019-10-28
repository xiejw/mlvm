package ir

import (
	"fmt"

	"github.com/xiejw/mlvm/mlvm/internal/naming"
)

// Creates a new Instruction in Func.
//
// This is convenient to use given it will try to auto generate the Instruction name and it will
// panic for any error.
func (f *Func) NewInstructionOrDie(op *Op, operands ...*Tensor) *Instruction {
	baseName := op.BaseName()
	var name string
	for {
		f.opNameIndex += 1
		name = naming.DefaultInstructionName(baseName, f.opNameIndex)
		if f.nameStore[name] == nil {
			break
		}
	}
	ins, err := f.NewInstructionWithName(name, op, operands...)
	if err != nil {
		panic(fmt.Errorf("Unexpected error during creating instruction for %v: %w", name, err))
	}
	return ins
}

func (f *Func) NewInstructionWithName(
	name string, op *Op, operands ...*Tensor,
) (*Instruction, error) {

	err := f.mustNotFreezed()
	if err != nil {
		return nil, err
	}

	ins, err := newInstruction(name, op, operands...)
	if err != nil {
		return nil, err
	}

	err = f.registerName(name, ins, true /* registerOnce */)
	if err != nil {
		return nil, err
	}

	f.instructions = append(f.instructions, ins)
	return ins, nil
}
