package code

import (
	"encoding/binary"
	"fmt"
)

type Opcode byte

const (
	// Opcode Name should be at most 10 chars.
	//
	OpConstant Opcode = iota // Loads constant object, uint16 index, from Program.
	OpLoadG                  // Loads object, uint16 index, from Global memory.
	OpStoreG                 // Stores objec, uint16 index, to Global memory.
	OpLoadT                  // Loads Tensor from Tensor store.
	OpStoreT                 // Loads Tensor from Tensor store.
	OpPrngNew                // Creates a new Prng source. The top stack operand is the seed.
	OpPrngDist               // Creates an Array with distribution (uint16 dist type index). Two stack operands are prng source (top), shape.
	OpTensor                 // Creates a new Tensor. Two stack operands are shape, array (top).
	OpAdd                    // Adds two stack operands.
)

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
	OpLoadG:    {"OpLoadG", []int{2}},
	OpStoreG:   {"OpStoreG", []int{2}},
	OpLoadT:    {"OpLoadT", []int{}},
	OpStoreT:   {"OpStoreT", []int{}},
	OpPrngNew:  {"OpPrngNew", []int{}},
	OpPrngDist: {"OpPrngDist", []int{2}},
	OpTensor:   {"OpTensor", []int{}},
	OpAdd:      {"OpAdd", []int{}},
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
