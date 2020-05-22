package vm

import (
	"fmt"

	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/vm/kernel"
)

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
