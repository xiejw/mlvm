package engine

import (
	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/vm/runtime"
)

type VM struct {
	instructions code.Instructions
	constants    []code.Object

	stack *runtime.Stack
}

func NewVM(program *code.Program) *VM {
	return &VM{
		instructions: program.Instructions,
		constants:    program.Constants,
		stack:        runtime.NewStack(),
	}
}

func (vm *VM) Run() error {
	return nil
}
