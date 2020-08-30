// Package `vm` provides the implementation of the stack virtual machine.
package vm

import (
	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/vm/mach"
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
	stack       *mach.Stack
	globalMem   *mach.Memory
	tensorStore mach.TensorStore
}

func NewVM(program *code.Program) *VM {
	return NewVMWithTensorStore(program, mach.NewTensorStore())
}

func NewVMWithTensorStore(program *code.Program, tensorStore mach.TensorStore) *VM {
	return &VM{
		instructions: program.Instructions,
		constants:    program.Constants,
		stack:        mach.NewStack(),
		globalMem:    mach.NewMemory(),
		tensorStore:  tensorStore,
	}
}

func (vm *VM) TensorStore() mach.TensorStore {
	return vm.tensorStore
}
