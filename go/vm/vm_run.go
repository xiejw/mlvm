package vm

import (
	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/object/prng64"
	"github.com/xiejw/mlvm/go/vm/kernel"
)

type Outputs []object.Object

// Run is expected to be call multiple times.
//
// Lifetime invariance.
//   - Upon and after run, the stack is empty.
//   - Upon run, the memory is reset.
//   - Across multiple runs, TensorStore is the only approach to persist data.
//
// During any run, if any error raised, the system's state becomes unknown. Creating a new VM is
// recommended.
func (vm *VM) Run() (Outputs, error) {
	end := len(vm.instructions)

	for ip := 0; ip < end; ip++ {

		op := code.Opcode(vm.instructions[ip])
		switch op {
		////////////////////////////////////////////////////////////////////////////////////////////////
		// Load/Stores (Constants, Global Memory, etc)
		case code.OpCONST:
			constantIndex := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip += 2

			if constantIndex >= len(vm.constants) {
				return nil, vm.canonicalError(op, "const (id: %v) does not exist.", constantIndex)
			}

			err := vm.stack.Push(vm.constants[constantIndex])
			if err != nil {
				return nil, vm.canonicalError(op, "internal error: %v.", err)
			}

		case code.OpSTORE:
			memSlotIndex := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip += 2

			o, err := vm.pop()
			if err != nil {
				return nil, vm.canonicalError(op, "expect to get object for store: %v.", err)
			}
			err = vm.globalMem.Set(memSlotIndex, o)
			if err != nil {
				return nil, vm.canonicalError(op,
					"failed to store object to global memory at %v: %v.", memSlotIndex, err)
			}

		case code.OpLOAD:
			memSlotIndex := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip += 2

			o, err := vm.globalMem.Get(memSlotIndex)
			if err != nil {
				return nil, vm.canonicalError(op,
					"failed to load object from global memory at %v: %v.", memSlotIndex, err)
			}

			err = vm.stack.Push(o)
			if err != nil {
				return nil, vm.canonicalError(op, "internal error: %v.", err)
			}

		case code.OpLOADS:
			key, err := vm.popString()
			if err != nil {
				return nil, vm.canonicalError(op, "expect to get key name from stack: %v.", err)
			}

			tensor, err := vm.tensorStore.Load(key.Value)
			if err != nil {
				return nil, vm.canonicalError(op,
					"failed to load tensor (key: \"%s\") from tensor store: %v.", key.Value, err)
			}

			err = vm.stack.Push(tensor)
			if err != nil {
				return nil, vm.canonicalError(op, "internal error: %v.", err)
			}

			////////////////////////////////////////////////////////////////////////////////////////////////
			// Prng
		case code.OpRNG:
			seed, err := vm.popInteger()
			if err != nil {
				return nil, vm.canonicalError(op, "expect to get Prng seed from stack: %v.", err)
			}

			prng := &object.Prng{Source: prng64.NewPrng64(uint64(seed.Value))}
			err = vm.stack.Push(prng)
			if err != nil {
				return nil, vm.canonicalError(op, "internal error: %v.", err)
			}

		case code.OpRNGT:
			distType := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2

			o, err := vm.pop()
			if err != nil {
				return nil, vm.canonicalError(op, "expect to get Prng object from stack: %v.", err)
			}
			prng, ok := o.(*object.Prng)
			if !ok {
				return nil, vm.canonicalError(op, "expect Prng source from stack: %v.", err)
			}

			shape, err := vm.popShape()
			if err != nil {
				return nil, vm.canonicalError(op, "expect shape from stack: %v.", err)
			}

			size := shape.Size()

			value := make([]float32, size)
			prng.FillDist(object.DistType(distType), value)

			err = vm.stack.Push(&object.Array{value})
			if err != nil {
				return nil, vm.canonicalError(op, "internal error: %v.", err)
			}

		////////////////////////////////////////////////////////////////////////////////////////////////
		// Tensor Related.
		case code.OpTensor:
			array, err := vm.popArray()
			if err != nil {
				return nil, err
			}

			shape, err := vm.popShape()
			if err != nil {
				return nil, err
			}

			tensor := &object.Tensor{shape, array}
			err = vm.stack.Push(tensor)
			if err != nil {
				return nil, vm.canonicalError(op, "internal error: %v.", err)
			}

		case code.OpTADD:
			operand1, err := vm.popTensor()
			if err != nil {
				return nil, err
			}
			operand2, err := vm.popTensor()
			if err != nil {
				return nil, err
			}

			tensor, err := kernel.TensorAdd(operand1, operand2)
			if err != nil {
				return nil, vm.canonicalError(op, "internal error: %v.", err)
			}

			err = vm.stack.Push(tensor)
			if err != nil {
				return nil, vm.canonicalError(op, "internal error: %v.", err)
			}

		default:
			startIndex := ip
			numInstructionsToPrint := 5
			return nil, vm.canonicalError(op, "unsupported Opcode in vm at @%d:\n\n%v\n",
				ip,
				code.Instructions(
					vm.instructions[ip:]).DebugString(startIndex, numInstructionsToPrint))
		}
	}
	return vm.popOutputs()
}

// Clears the stack and moves items (in reverse order) as outputs.
func (vm *VM) popOutputs() (Outputs, error) {
	var outputs Outputs
	stack := vm.stack

	for {
		if stack.Top() == nil {
			break
		}

		item, err := stack.Pop()
		if err != nil {
			return nil, err
		}

		outputs = append(outputs, item)
	}
	return outputs, nil
}
