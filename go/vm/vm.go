package vm

import (
	"fmt"

	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
)

type VM struct {
	instructions code.Instructions
	constants    []object.Object

	stack *Stack
}

func NewVM(program *code.Program) *VM {
	return &VM{
		instructions: program.Instructions,
		constants:    program.Constants,
		stack:        NewStack(),
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
				return fmt.Errorf("program error: Opcode: %v: const (id: %v) does not exist", op, constIndex)
			}

			err := vm.stack.Push(vm.constants[constIndex])
			if err != nil {
				return fmt.Errorf("program error: Opcode: %v: internal error: %w", op, err)
			}
			ip += 2
		case code.OpTensor:
			arrayObject, err := vm.stack.Pop()
			if err != nil {
				return fmt.Errorf("program error: Opcode: %v: failed to pop array from stack: %v", op, err)
			}
			array, ok := arrayObject.(*object.Array)
			if !ok {
				return fmt.Errorf("program error: Opcode: %v: failed to pop array from stack: wrong type", op)
			}

			shapeObject, err := vm.stack.Pop()
			if err != nil {
				return fmt.Errorf("program error: Opcode: %v: failed to pop shape from stack: %v", op, err)
			}
			shape, ok := shapeObject.(*object.Shape)
			if !ok {
				return fmt.Errorf("program error: Opcode: %v: failed to pop shape from stack: wrong type", op)
			}

			tensor := &object.Tensor{shape, array}
			err = vm.stack.Push(tensor)
			if err != nil {
				return fmt.Errorf("program error: Opcode: %v: internal error: %w", op, err)
			}

		default:
			return fmt.Errorf("program error: Opcode: %v: got unsupported Opcode at @%5d", op, ip)
		}
		ip++

	}

	return nil
}

func (vm *VM) StackTop() object.Object {
	return vm.stack.Top()
}
