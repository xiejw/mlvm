package vm

import (
	"fmt"

	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/vm/kernel"
)

type VM struct {
	instructions code.Instructions
	data         []object.Object

	stack *Stack
}

func NewVM(program *code.Program) *VM {
	return &VM{
		instructions: program.Instructions,
		data:         program.Data,
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
		case code.OpData:
			constIndex := int(code.ReadUint16(vm.instructions[ip+1:]))
			if constIndex >= len(vm.data) {
				return fmt.Errorf("program error: Opcode: %v: const (id: %v) does not exist", op, constIndex)
			}

			err := vm.stack.Push(vm.data[constIndex])
			if err != nil {
				return fmt.Errorf("program error: Opcode: %v: internal error: %w", op, err)
			}
			ip += 2

		case code.OpTensor:
			array, err := vm.popArray(op)
			if err != nil {
				return err
			}

			shape, err := vm.popShape(op)
			if err != nil {
				return err
			}

			tensor := &object.Tensor{shape, array}
			err = vm.stack.Push(tensor)
			if err != nil {
				return fmt.Errorf("program error: Opcode: %v: internal error: %w", op, err)
			}

		case code.OpAdd:
			operand1, err := vm.popTensor(op)
			if err != nil {
				return err
			}
			operand2, err := vm.popTensor(op)
			if err != nil {
				return err
			}

			fmt.Printf("%v\n", operand1)
			fmt.Printf("%v\n", operand2)

			tensor, err := kernel.TensorAdd(operand1, operand2)
			if err != nil {
				return fmt.Errorf("program error: Opcode: %v: internal error: %w", op, err)
			}

			err = vm.stack.Push(tensor)
			if err != nil {
				return fmt.Errorf("program error: Opcode: %v: internal error: %w", op, err)
			}

		default:
			return fmt.Errorf("program error: Opcode: `%v`: unsupported Opcode in vm at @%5d", op, ip)
		}
		ip++

	}

	return nil
}

func (vm *VM) StackTop() object.Object {
	return vm.stack.Top()
}

func (vm *VM) popArray(op code.Opcode) (*object.Array, error) {
	arrayObject, err := vm.stack.Pop()
	if err != nil {
		return nil, fmt.Errorf("program error: Opcode: %v: failed to pop array from stack: %v", op, err)
	}
	array, ok := arrayObject.(*object.Array)
	if !ok {
		return nil, fmt.Errorf("program error: Opcode: %v: failed to pop array from stack: wrong type", op)
	}
	return array, nil
}

func (vm *VM) popShape(op code.Opcode) (*object.Shape, error) {
	shapeObject, err := vm.stack.Pop()
	if err != nil {
		return nil, fmt.Errorf("program error: Opcode: %v: failed to pop shape from stack: %v", op, err)
	}
	shape, ok := shapeObject.(*object.Shape)
	if !ok {
		return nil, fmt.Errorf("program error: Opcode: %v: failed to pop shape from stack: wrong type", op)
	}
	return shape, nil
}

func (vm *VM) popTensor(op code.Opcode) (*object.Tensor, error) {
	tensorObject, err := vm.stack.Pop()
	if err != nil {
		return nil, fmt.Errorf("program error: Opcode: %v: failed to pop tensor from stack: %v", op, err)
	}
	tensor, ok := tensorObject.(*object.Tensor)
	if !ok {
		return nil, fmt.Errorf("program error: Opcode: %v: failed to pop tensor from stack: wrong type", op)
	}
	return tensor, nil
}
