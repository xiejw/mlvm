package vm

import (
	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
)

type VM struct {
	// Copied from Program.
	instructions code.Instructions
	data         []object.Object

	// Internal States.
	stack     *Stack
	globalMem *Memory
}

func NewVM(program *code.Program) *VM {
	return &VM{
		instructions: program.Instructions,
		data:         program.Data,
		stack:        NewStack(),
		globalMem:    NewMemory(),
	}
}

func (vm *VM) StackTop() object.Object {
	return vm.stack.Top()
}
