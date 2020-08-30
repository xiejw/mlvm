package code

import (
	"encoding/binary"
	"fmt"
)

type Opcode byte

const (
	// Opcode string name should be at most 10 chars. Enforced in test.
	//
	// For meaning, see Definition.
	OpCONST Opcode = iota
	OpPOP
	OpLOAD
	OpSTORE
	OpLOADS
	OpSTORES
	OpRNG
	OpRNGT
	OpRNGS
	OpT
	OpTADD
)

// Defines the string name and operand requirements.
type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	// Loads constant object from Program into stack.
	//
	// Operand: (uint16) object index.
	// Stack  : push the object to the top.
	OpCONST: {"OpCONST", []int{2}},

	// Pops out top item on stack.
	//
	// Operand: no.
	// Stack  : pop the object from the top.
	OpPOP: {"OpPOP", []int{}},

	// Loads top object from global memory.
	//
	// Operand: (uint16) object index.
	// Stack  : push the object to the top.
	OpLOAD: {"OpLOAD", []int{2}},

	// Stores object into global memory.
	//
	// Operand: (uint16) object index.
	// Stack  : pop the object from the top.
	OpSTORE: {"OpSTORE", []int{2}},

	// Loads object from key-value store.
	//
	// Operand: no.
	// Stack  : pops the top item and uses it as string key.
	OpLOADS: {"OpLOADS", []int{}},

	// Stores the object into key-value store.
	//
	// Operand: no.
	// Stack  : pops the top item and uses it as the object.
	//          pops the second item and uses it as (String) key.
	OpSTORES: {"OpSTORES", []int{}},

	// Creates a new rng source.
	//
	// Operand: no.
	// Stack  : pops the top item and uses it as (Integer) seed.
	//          stores the rng source into the stack.
	OpRNG: {"OpRNG", []int{}},

	// Creates a Tensor with selectd distribution.
	//
	// Operand: (uint16) distribution type index.
	// Stack  : pops the top item and uses it as rng source.
	//          pops the second item and uses it as (Shape).
	//          stores the Tensor into the stack.
	OpRNGT: {"OpRNGT", []int{2}},

	// Splits the rng source.
	//
	// Operand: no
	// Stack  : pops the top item and uses it as rng source.
	//          stores the first item of the result into the stack.
	//          then stores the second item of the result into the stack.
	OpRNGS: {"OpRNGS", []int{}},

	// Creates a new Tensor.
	//
	// Operand: no.
	// Stack  : pops the top item and uses it as (Array).
	//          pops the second item and uses it as (Shape).
	//          stores the Tensor into the stack.
	OpT: {"OpT", []int{}},

	// Adds two tensors.
	//
	// Operand: no.
	// Stack  : pops the top item and uses it as second operand.
	//          then pops the top item and uses it as first operand.
	//          stores the result into the stack.
	OpTADD: {"OpTADD", []int{}},
}

// Looks up the Definition of the op code.
func Lookup(op Opcode) (*Definition, error) {
	def, ok := definitions[op]
	if !ok {
		return nil, fmt.Errorf("Opcode %d undefined", op)
	}

	return def, nil
}

// Encodes the op code with its operands, if any.
//
// This is Opcode encoder. The decoder is in vm.
func MakeOp(op Opcode, operands ...int) ([]byte, error) {
	def, err := Lookup(op)
	if err != nil {
		return nil, err
	}

	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	if len(def.OperandWidths) != len(operands) {
		return nil, fmt.Errorf("Operand `%v` counts mismatch: expected %v, got %v",
			op, len(def.OperandWidths), len(operands))
	}

	offset := 1
	for i, o := range operands {
		w := def.OperandWidths[i]
		switch w {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		default:
			panic(fmt.Sprintf("unsupported width (%v) for op: %v", w, op))
		}
	}

	return instruction, nil
}

////////////////////////////////////////////////////////////////////////////////
// Public Helper Methods
////////////////////////////////////////////////////////////////////////////////

func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}

////////////////////////////////////////////////////////////////////////////////
// String Related.
////////////////////////////////////////////////////////////////////////////////

func (opcode Opcode) String() string {
	def, ok := definitions[opcode]
	if !ok {
		panic(fmt.Errorf("Opcode %d undefined", opcode))
	}
	return def.Name
}
