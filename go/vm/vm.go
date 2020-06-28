package vm

import (
	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
)

type VM struct {
	// Copied from Program. Shall not mutate.
	instructions code.Instructions
	constants    []object.Object

	// Internal States.
	stack       *Stack
	globalMem   *Memory
	tensorStore TensorStore
}

func NewVM(program *code.Program) *VM {
	return &VM{
		instructions: program.Instructions,
		constants:    program.Constants,
		stack:        NewStack(),
		globalMem:    NewMemory(),
		tensorStore:  NewTensorStore(),
	}
}

func (vm *VM) ResetTensorStore(tensorStore TensorStore) TensorStore {
	oldTensorStore := vm.tensorStore
	vm.tensorStore = tensorStore
	return oldTensorStore
}

func (vm *VM) StackTop() object.Object {
	return vm.stack.Top()
}
