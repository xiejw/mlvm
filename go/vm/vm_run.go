package vm

import (
	"log"

	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/vm/kernel"
	"github.com/xiejw/mlvm/go/vm/prng64"
)

const vmErr = "virtual machine error: current opcode %v"

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
func (vm *VM) Run() (Outputs, *errors.DError) {
	end := len(vm.instructions)

	for ip := 0; ip < end; ip++ {

		op := code.Opcode(vm.instructions[ip])
		switch op {
		////////////////////////////////////////////////////////////////////////////////////////////////
		// Load/Stores (Constants, Global Memory, etc)
		case code.OpCONST:
			cIndex := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip += 2

			if cIndex >= len(vm.constants) {
				return nil, errors.New("const (id: %v) not exist.", cIndex).EmitNote(vmErr, op)
			}

			err := vm.stack.Push(vm.constants[cIndex])
			if err != nil {
				return nil, err.EmitNote("failed to push const to stack.").EmitNote(vmErr, op)
			}

		case code.OpSTORE:
			mIndex := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip += 2

			o, err := vm.pop()
			if err != nil {
				return nil, err.EmitNote("failed to get object from stack.").EmitNote(vmErr, op)
			}
			err = vm.globalMem.Set(mIndex, o)
			if err != nil {
				return nil, err.EmitNote("failed to store obj to mem @%v.", mIndex).EmitNote(vmErr, op)
			}

		case code.OpLOAD:
			mIndex := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip += 2

			o, err := vm.globalMem.Get(mIndex)
			if err != nil {
				return nil, err.EmitNote("failed to load obj from mem @%v.", mIndex).EmitNote(vmErr, op)
			}

			err = vm.stack.Push(o)
			if err != nil {
				return nil, err.EmitNote("failed to push to stack.").EmitNote(vmErr, op)
			}

		case code.OpLOADS:
			key, err := vm.popString()
			if err != nil {
				return nil, err.EmitNote("failed to get key from stack.").EmitNote(vmErr, op)
			}

			keyStr := key.Value
			tensor, err := vm.store.Load(keyStr)
			if err != nil {
				return nil, err.EmitNote("failed to load obj (k:\"%s\") from store.", keyStr).EmitNote(vmErr, op)
			}

			err = vm.stack.Push(tensor)
			if err != nil {
				return nil, err.EmitNote("failed to push to stack.").EmitNote(vmErr, op)
			}

		case code.OpIOR:
			n := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip += 2

			log.Printf("vm infeed %v objects start", n)

			for i := 0; i < n; i++ {
				o := <-vm.c
				err := vm.stack.Push(o)
				if err != nil {
					return nil, err.EmitNote("failed to push to stack.").EmitNote(vmErr, op)
				}
			}

			log.Printf("vm infeed %v objects end", n)

			////////////////////////////////////////////////////////////////////////////////////////////////
			// Rng
		case code.OpRNG:
			seed, err := vm.popInteger()
			if err != nil {
				return nil, err.EmitNote("failed to get rng seed from stack.").EmitNote(vmErr, op)
			}
			src := prng64.NewPrng64(uint64(seed.Value))
			prng := &object.Rng{src.Seed, src.Gamma, src.NextGammaSeed}
			err = vm.stack.Push(prng)
			if err != nil {
				return nil, err.EmitNote("failed to push to stack.").EmitNote(vmErr, op)
			}

		case code.OpRNGT:
			distType := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2

			o, err := vm.pop()
			if err != nil {
				return nil, err.EmitNote("failed to get rng from stack.").EmitNote(vmErr, op)
			}
			rng, ok := o.(*object.Rng)
			if !ok {
				return nil, errors.New("failed to cast obj to rng.").EmitNote(vmErr, op)
			}

			shape, err := vm.popShape()
			if err != nil {
				return nil, err.EmitNote("failed to get hape from stack.").EmitNote(vmErr, op)
			}

			size := shape.Size()

			value := make([]float32, size)
			prng := prng64.Prng64{
				Seed:          rng.Seed,
				Gamma:         rng.Gamma,
				NextGammaSeed: rng.NextGammaSeed,
			}

			prng64.FillDist(&prng, prng64.DistType(distType), value)

			err = vm.stack.Push(&object.Array{value})
			if err != nil {
				return nil, err.EmitNote("failed to push to stack.").EmitNote(vmErr, op)
			}

		////////////////////////////////////////////////////////////////////////////////////////////////
		// Tensor Related.
		case code.OpT:
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
				return nil, err.EmitNote("failed to push to stack.").EmitNote(vmErr, op)
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

			tensor, err1 := kernel.TensorAdd(operand1, operand2)
			if err1 != nil {
				return nil, errors.From(err1).EmitNote("kernel error for TensorAdd").EmitNote(vmErr, op)
			}

			err = vm.stack.Push(tensor)
			if err != nil {
				return nil, err.EmitNote("failed to push to stack.").EmitNote(vmErr, op)
			}

		default:
			startIndex := ip
			numInstructionsToPrint := 5
			return nil, errors.New("unsupported OpCode in vm at @%d:\n\n%v\n",
				ip,
				code.Instructions(
					vm.instructions[ip:]).DebugString(startIndex, numInstructionsToPrint))
		}
	}
	return vm.popOutputs()
}

// Clears the stack and moves items (in reverse order) as outputs.
func (vm *VM) popOutputs() (Outputs, *errors.DError) {
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
