package vm

import (
	"fmt"

	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/object/prng64"
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
		////////////////////////////////////////////////////////////////////////////////////////////////
		// Load/Stores (Constants, Global Memory, etc)
		case code.OpConstant:
			constantIndex := int(code.ReadUint16(vm.instructions[ip+1:]))
			if constantIndex >= len(vm.constants) {
				return vm.canonicalError(op, "const (id: %v) does not exist.", constantIndex)
			}

			err := vm.stack.Push(vm.constants[constantIndex])
			if err != nil {
				return vm.canonicalError(op, "internal error: %v.", err)
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
					"failed to store object to global memory at %v: %v.", memSlotIndex, err)
			}
			ip += 2

			////////////////////////////////////////////////////////////////////////////////////////////////
			// Prng
		case code.OpPrngNew:
			seed, err := vm.popInteger(op)
			if err != nil {
				return vm.canonicalError(op, "failed to get Prng seed from stack: %v.", err)
			}

			prng := prng64.NewPrng64(uint64(seed.Value))
			err = vm.stack.Push(prng)
			if err != nil {
				return vm.canonicalError(op, "internal error: %v.", err)
			}

		case code.OpPrngDist:
			distType := code.ReadUint16(vm.instructions[ip+1:])

			o, err := vm.pop(op)
			if err != nil {
				return vm.canonicalError(op, "failed to get Prng source from stack: %v.", err)
			}
			prng, ok := o.(*prng64.Prng64)
			if !ok {
				return vm.canonicalError(op, "expect Prng source from stack: %v.", err)
			}

			shape, err := vm.popShape(op)
			if err != nil {
				return vm.canonicalError(op, "failed to get shape for OpPrngDist: %v.", err)
			}

			size := shape.Size()

			value := make([]float32, size)
			prng.FillDist(prng64.DistType(distType), value)

			err = vm.stack.Push(&object.Array{value})
			if err != nil {
				return vm.canonicalError(op, "internal error: %v.", err)
			}

		////////////////////////////////////////////////////////////////////////////////////////////////
		// Tensor Related.
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
				return vm.canonicalError(op, "internal error: %v.", err)
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
				return vm.canonicalError(op, "internal error: %v.", err)
			}

			err = vm.stack.Push(tensor)
			if err != nil {
				return vm.canonicalError(op, "internal error: %v.", err)
			}

		default:
			return vm.canonicalError(op, "unsupported Opcode in vm at @%5d.", ip)
		}
		ip++

	}

	return nil
}
