// Package `mach` provides the implementation of the stack virtual machine.
package mach

import (
	"github.com/xiejw/mlvm/vm/code"
	"github.com/xiejw/mlvm/vm/mach/parts"
	"github.com/xiejw/mlvm/vm/object"
)

// VM is a machine which runs the user provided program. The program is immutable; while the
// key-value store might be mutated according to the program.
//
// Check Run() to see the lifetime of VM.
type VM struct {
	// Copied from Program. Immutate.
	instructions code.Instructions
	constants    []object.Object

	// Internal States.
	stack     *parts.Stack
	globalMem *parts.Memory
	store     parts.TensorStore
	c         chan object.Object
}

func NewVM(program *code.Program) *VM {
	return &VM{
		instructions: program.Instructions,
		constants:    program.Constants,
		stack:        parts.NewStack(),
		globalMem:    parts.NewMemory(),
		store:        parts.NewTensorStore(),
		c:            make(chan object.Object),
	}
}

// Returns the infeed chan to enqueue objects.
func (m *VM) InfeedChan() chan<- object.Object {
	return m.c
}
