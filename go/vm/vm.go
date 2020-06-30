package vm

import (
	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
)

// VM is a machine which runs the user provided program. The program is not mutable; while the
// TensorStore might be mutated according to the program.
//
// Check Run() to see the lifetime of VM.
type VM struct {
	// Copied from Program. Immutate.
	instructions code.Instructions
	constants    []object.Object

	// Internal States.
	stack       *Stack
	globalMem   *Memory
	tensorStore TensorStore
}

func NewVM(program *code.Program) *VM {
	return NewVMWithTensorStore(program, NewTensorStore())
}

func NewVMWithTensorStore(program *code.Program, tensorStore TensorStore) *VM {
	return &VM{
		instructions: program.Instructions,
		constants:    program.Constants,
		stack:        NewStack(),
		globalMem:    NewMemory(),
		tensorStore:  tensorStore,
	}
}

func (vm *VM) TensorStore() TensorStore {
	return vm.tensorStore
}
