package mach

import (
	"log"

	"github.com/xiejw/mlvm/vm/base/errors"
	"github.com/xiejw/mlvm/vm/code"
	"github.com/xiejw/mlvm/vm/mach/mat"
	"github.com/xiejw/mlvm/vm/mach/prng64"
	"github.com/xiejw/mlvm/vm/mach/tensorarray"
	"github.com/xiejw/mlvm/vm/object"
)

const vmErr = "virtual machine error: current opcode %v"

type Outputs []object.Object

// Run is expected to be call multiple times.
//
// # Lifetime invariance.
//
//   - Upon and after run, the stack is empty.
//   - Upon run, the memory is reset. Though toward optimization, reset might not be performed.
//   - Across multiple runs, key value store is the only approach to persist data.
//
//
// # Internal Representations: Tensor vs TensorArray
//
// Inside vm, TensorArray is the source of truth for tensor operations. So, conversion is performed
// at all boundaries, including loading/storing from key value store, returning outputs, infeeding,
// load constants, etc.
//
//
// # Error Handling.
//
// During any run, if any error raised, the system's state becomes unknown. Creating a new VM is
// recommended.
func (vm *VM) Run() (Outputs, error) {
	end := len(vm.instructions)

	for ip := 0; ip < end; ip++ {

		op := code.Opcode(vm.instructions[ip])
		switch op {

		////////////////////////////////////////////////////////////////////////////
		// Load/Store/Move (Constants, Global Memory, etc)
		////////////////////////////////////////////////////////////////////////////

		case code.OpCONST:
			cIndex := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip += 2

			if cIndex >= len(vm.constants) {
				return nil, errors.New("const (id: %v) not exist.", cIndex).EmitNote(vmErr, op)
			}
			c := vm.constants[cIndex]
			c = normalizeObject(c)
			err := vm.stack.Push(c)
			if err != nil {
				return nil, errors.From(err).EmitNote("failed to push const to stack.").EmitNote(vmErr, op)
			}

		case code.OpSTORE:
			mIndex := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip += 2

			o, err := vm.pop()
			if err != nil {
				return nil, errors.From(err).EmitNote("failed to get object from stack.").EmitNote(vmErr, op)
			}
			err = vm.globalMem.Set(mIndex, o)
			if err != nil {
				return nil, errors.From(err).EmitNote("failed to store obj to mem @%v.", mIndex).EmitNote(vmErr, op)
			}

		case code.OpLOAD:
			mIndex := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip += 2

			o, err := vm.globalMem.Get(mIndex)
			if err != nil {
				return nil, errors.From(err).EmitNote("failed to load obj from mem @%v.", mIndex).EmitNote(vmErr, op)
			}

			err = vm.stack.Push(o)
			if err != nil {
				return nil, errors.From(err).EmitNote("failed to push to stack.").EmitNote(vmErr, op)
			}

		case code.OpMOVE:
			mIndex := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip += 2

			o, err := vm.globalMem.Drop(mIndex)
			if err != nil {
				return nil, errors.From(err).EmitNote("failed to obtain obj from mem @%v.", mIndex).EmitNote(vmErr, op)
			}

			err = vm.stack.Push(o)
			if err != nil {
				return nil, errors.From(err).EmitNote("failed to push to stack.").EmitNote(vmErr, op)
			}

		////////////////////////////////////////////////////////////////////////////
		// External I/Os. Key-value Store and Infeed Channel.
		////////////////////////////////////////////////////////////////////////////

		case code.OpLOADS:
			key, err := vm.popString()
			if err != nil {
				return nil, errors.From(err).EmitNote("failed to get key from stack.").EmitNote(vmErr, op)
			}

			keyStr := key.Value
			tensor, err := vm.store.Load(keyStr)
			if err != nil {
				return nil, errors.From(err).EmitNote("failed to load obj (k:\"%s\") from store.", keyStr).EmitNote(vmErr, op)
			}

			err = vm.stack.Push(normalizeObject(tensor))
			if err != nil {
				return nil, errors.From(err).EmitNote("failed to push to stack.").EmitNote(vmErr, op)
			}

		case code.OpIOR:
			n := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip += 2

			log.Printf("vm infeed %v objects start", n)

			for i := 0; i < n; i++ {
				o := <-vm.c
				o = normalizeObject(o)
				err := vm.stack.Push(o)
				if err != nil {
					return nil, errors.From(err).EmitNote("failed to push to stack.").EmitNote(vmErr, op)
				}
			}

			log.Printf("vm infeed %v objects end", n)

		////////////////////////////////////////////////////////////////////////////
		// Random Number Generator.
		////////////////////////////////////////////////////////////////////////////

		case code.OpRNG:
			seed, err := vm.popInteger()
			if err != nil {
				return nil, errors.From(err).EmitNote("failed to get rng seed from stack.").EmitNote(vmErr, op)
			}
			src := prng64.NewPrng64(uint64(seed.Value))
			prng := &object.Rng{src.Seed, src.Gamma, src.NextGammaSeed}
			err = vm.stack.Push(prng)
			if err != nil {
				return nil, errors.From(err).EmitNote("failed to push to stack.").EmitNote(vmErr, op)
			}

		case code.OpRNGT:
			distType := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2

			o, err := vm.pop()
			if err != nil {
				return nil, errors.From(err).EmitNote("failed to get rng from stack.").EmitNote(vmErr, op)
			}
			rng, ok := o.(*object.Rng)
			if !ok {
				return nil, errors.New("failed to cast obj to rng.").EmitNote(vmErr, op)
			}

			shape, err := vm.popShape()
			if err != nil {
				return nil, errors.From(err).EmitNote("failed to get hape from stack.").EmitNote(vmErr, op)
			}

			size := shape.Size()

			value := make([]float32, size)
			prng := prng64.Prng64{
				Seed:          rng.Seed,
				Gamma:         rng.Gamma,
				NextGammaSeed: rng.NextGammaSeed,
			}

			prng64.FillDist(&prng, prng64.DistType(distType), value)

			err = vm.stack.Push(tensorarray.FromRaw(shape.Dims, value))
			if err != nil {
				return nil, errors.From(err).EmitNote("failed to push to stack.").EmitNote(vmErr, op)
			}

		////////////////////////////////////////////////////////////////////////////
		// Tensor Related.
		////////////////////////////////////////////////////////////////////////////

		case code.OpT:
			array, err := vm.popArray()
			if err != nil {
				return nil, err
			}

			shape, err := vm.popShape()
			if err != nil {
				return nil, err
			}

			tensor := tensorarray.FromRaw(shape.Dims, array.Value)
			err = vm.stack.Push(tensor)
			if err != nil {
				return nil, errors.From(err).EmitNote("failed to push to stack.").EmitNote(vmErr, op)
			}

		case code.OpTSHAPE:
			ta, err := vm.popTensor()
			if err != nil {
				return nil, errors.From(err).EmitNote("failed to pop tensor from stack").EmitNote(vmErr, op)
			}

			shape := object.NewShape(ta.Dims)
			err = vm.stack.Push(shape)
			if err != nil {
				return nil, errors.From(err).EmitNote("failed to push to stack.").EmitNote(vmErr, op)
			}

		case code.OpTADD:
			fallthrough
		case code.OpTMINUS:
			fallthrough
		case code.OpTMUL:
			op_type := mat.BinaryOpType(int(op) - int(code.OpTADD) + int(mat.BinaryAdd))

			lhs, rhs, err := vm.popTwoTensorsInSeq()
			if err != nil {
				return nil, errors.EmitNote(err, vmErr, op)
			}

			tensor, err := mat.BinaryOp(lhs, rhs, op_type)
			if err != nil {
				return nil, errors.From(err).EmitNote("unexpected error for binary op").EmitNote(vmErr, op)
			}

			err = vm.stack.Push(tensor)
			if err != nil {
				return nil, errors.From(err).EmitNote("failed to push to stack.").EmitNote(vmErr, op)
			}

		case code.OpTREDUCE:
			merge_type := mat.MergeType(code.ReadUint16(vm.instructions[ip+1:]))
			ip += 2

			ta, err := vm.popTensor()
			if err != nil {
				return nil, err
			}

			r, err := mat.Reduce(ta, merge_type)
			if err != nil {
				return nil, errors.From(err).EmitNote("unexpected error for reduce").EmitNote(vmErr, op)
			}

			err = vm.stack.Push(r)
			if err != nil {
				return nil, errors.From(err).EmitNote("failed to push to stack.").EmitNote(vmErr, op)
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

///////////////////////////////////////////////////////////////////////////////
// Helper Methods.
///////////////////////////////////////////////////////////////////////////////

func (vm *VM) popTwoTensorsInSeq() (*tensorarray.TensorArray, *tensorarray.TensorArray, error) {

	rhs, err := vm.popTensor()
	if err != nil {
		return nil, nil, errors.EmitNote(err, "failed to the right hand side operand.")
	}
	lhs, err := vm.popTensor()
	if err != nil {
		return nil, nil, errors.EmitNote(err, "failed to the left hand side operand.")
	}
	return lhs, rhs, nil
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

		outputs = append(outputs, toStandardObject(item))
	}
	return outputs, nil
}

// Normalize external Object into internal representations.
func normalizeObject(o object.Object) object.Object {
	if o.Type() != object.TensorType {
		return o
	}

	t, ok := o.(*object.Tensor)
	if !ok {
		panic("external representation for tensor should be tensor.")
	}

	return tensorarray.FromTensor(t)

}

// Covnerts internal representations of Object to external standard versions.
func toStandardObject(o object.Object) object.Object {
	if o.Type() != object.TensorType {
		return o
	}

	ta, ok := o.(*tensorarray.TensorArray)
	if !ok {
		panic("internal representation for tensor should be tensor array.")
	}

	return ta.ToTensor()
}
