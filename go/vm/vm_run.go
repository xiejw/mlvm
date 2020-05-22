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
			dataIndex := int(code.ReadUint16(vm.instructions[ip+1:]))
			if dataIndex >= len(vm.data) {
				return vm.canonicalError(op, "const (id: %v) does not exist", dataIndex)
			}

			err := vm.stack.Push(vm.data[dataIndex])
			if err != nil {
				return vm.canonicalError(op, "internal error: %v", err)
			}
			ip += 2

		case code.OpStoreG:
			memSlotIndex := int(code.ReadUint16(vm.instructions[ip+1:]))
			o, err := vm.pop(op)
			if err != nil {
				return vm.canonicalError(op, "failed to get object for store: %v.", err)
			}
			err = vm.globalMem.Set(memSlotIndex, o)
			if err != nil {
				return vm.canonicalError(op,
					"failed to store object to global memory at %v: %v", memSlotIndex, err)
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
				return vm.canonicalError(op, "internal error: %v", err)
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
				return vm.canonicalError(op, "internal error: %v", err)
			}

			err = vm.stack.Push(tensor)
			if err != nil {
				return vm.canonicalError(op, "internal error: %v", err)
			}

		default:
			return vm.canonicalError(op, ": unsupported Opcode in vm at @%5d", ip)
		}
		ip++

	}

	return nil
}
