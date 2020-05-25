package vm

import (
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
			ip += 2

			if constantIndex >= len(vm.constants) {
				return vm.canonicalError(op, "const (id: %v) does not exist.", constantIndex)
			}

			err := vm.stack.Push(vm.constants[constantIndex])
			if err != nil {
				return vm.canonicalError(op, "internal error: %v.", err)
			}

		case code.OpStoreG:
			memSlotIndex := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip += 2

			o, err := vm.pop()
			if err != nil {
				return vm.canonicalError(op, "expect to get object for store: %v.", err)
			}
			err = vm.globalMem.Set(memSlotIndex, o)
			if err != nil {
				return vm.canonicalError(op,
					"failed to store object to global memory at %v: %v.", memSlotIndex, err)
			}

		case code.OpLoadG:
			memSlotIndex := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip += 2

			o, err := vm.globalMem.Get(memSlotIndex)
			if err != nil {
				return vm.canonicalError(op,
					"failed to load object from global memory at %v: %v.", memSlotIndex, err)
			}

			err = vm.stack.Push(o)
			if err != nil {
				return vm.canonicalError(op, "internal error: %v.", err)
			}

			////////////////////////////////////////////////////////////////////////////////////////////////
			// Prng
		case code.OpPrngNew:
			seed, err := vm.popInteger()
			if err != nil {
				return vm.canonicalError(op, "expect to get Prng seed from stack: %v.", err)
			}

			prng := prng64.NewPrng64(uint64(seed.Value))
			err = vm.stack.Push(prng)
			if err != nil {
				return vm.canonicalError(op, "internal error: %v.", err)
			}

		case code.OpPrngDist:
			distType := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2

			o, err := vm.pop()
			if err != nil {
				return vm.canonicalError(op, "expect to get Prng object from stack: %v.", err)
			}
			prng, ok := o.(*prng64.Prng64)
			if !ok {
				return vm.canonicalError(op, "expect Prng source from stack: %v.", err)
			}

			shape, err := vm.popShape()
			if err != nil {
				return vm.canonicalError(op, "expect shape from stack: %v.", err)
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
			array, err := vm.popArray()
			if err != nil {
				return err
			}

			shape, err := vm.popShape()
			if err != nil {
				return err
			}

			tensor := &object.Tensor{shape, array}
			err = vm.stack.Push(tensor)
			if err != nil {
				return vm.canonicalError(op, "internal error: %v.", err)
			}

		case code.OpAdd:
			operand1, err := vm.popTensor()
			if err != nil {
				return err
			}
			operand2, err := vm.popTensor()
			if err != nil {
				return err
			}

			tensor, err := kernel.TensorAdd(operand1, operand2)
			if err != nil {
				return vm.canonicalError(op, "internal error: %v.", err)
			}

			err = vm.stack.Push(tensor)
			if err != nil {
				return vm.canonicalError(op, "internal error: %v.", err)
			}

		default:
			end := ip + 5
			if end > len(vm.instructions) {
				end = len(vm.instructions)
			}
			return vm.canonicalError(op, "unsupported Opcode in vm at @%5d: %v",
				ip, code.Instructions(vm.instructions[ip:end]).DebugString(ip /*startIndex*/))
		}
		ip++

	}
	return nil
}
