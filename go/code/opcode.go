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
	OpLoadG
	OpStoreG
	OpLoadT
	OpStoreT
	OpPrngNew
	OpPrngDist
	OpTensor
	OpAdd
)

// Defines the string name and operand requirements.
type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	// Loads constant object, uint16 index, from Program.
	OpCONST: {"OpCONST", []int{2}},
	// Pops out top item on stack.
	OpPOP: {"OpPOP", []int{}},
	// Loads object, uint16 index, from Global memory.
	OpLoadG: {"OpLoadG", []int{2}},
	// Stores objec, uint16 index, to Global memory.
	OpStoreG: {"OpStoreG", []int{2}},
	// Loads Tensor from Tensor store.
	OpLoadT: {"OpLoadT", []int{}},
	// Loads Tensor from Tensor store.
	OpStoreT: {"OpStoreT", []int{}},
	// Creates a new Prng source. The top stack operand is the seed.
	OpPrngNew: {"OpPrngNew", []int{}},
	// Creates an Array with distribution (uint16 dist type index). Two stack operands are prng source
	// (top), shape.
	OpPrngDist: {"OpPrngDist", []int{2}},
	// Creates a new Tensor. Two stack operands are shape, array (top).
	OpTensor: {"OpTensor", []int{}},
	// Adds two stack operands.
	OpAdd: {"OpAdd", []int{}},
}

func Lookup(op Opcode) (*Definition, error) {
	def, ok := definitions[op]
	if !ok {
		return nil, fmt.Errorf("Opcode %d undefined", op)
	}

	return def, nil
}

func (opcode Opcode) String() string {
	def, ok := definitions[opcode]
	if !ok {
		panic(fmt.Errorf("Opcode %d undefined", opcode))
	}
	return def.Name
}

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

////////////////////////////////////////////////////////////////////////////////////////////////////
// Public Helper Methods

func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}
