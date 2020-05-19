package engine

import (
	"fmt"

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

	ip := 0
	end := len(vm.instructions)

	for {

		if ip >= end {
			break
		}

		op := code.Opcode(vm.instructions[ip])
		switch op {
		case code.OpConstant:
			constIndex := int(code.ReadUint16(vm.instructions[ip+1:]))
			if constIndex >= len(vm.constants) {
				return fmt.Errorf("program error: const (id: %v) does not exist", constIndex)
			}

			err := vm.stack.Push(vm.constants[constIndex])
			if err != nil {
				return fmt.Errorf("vm internal error: %w", err)
			}
			ip += 2
		default:
			return fmt.Errorf("got unsupported op at @%d: %v", ip, op)
		}
		ip++

	}

	return nil
}
