package code

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Instructions []byte

type Opcode byte

const (
	// Opcode Name should be at most 10 chars.
	//
	OpConstant Opcode = iota // Loads constant object, uint16 index, from Program.
	OpLoadG                  // Loads object, uint16 index, from Global memory.
	OpStoreG                 // Stores objec, uint16 index, to Global memory.
	OpPrngNew                // Creates a new Prng source. The top stack operand is the seed.
	OpPrngDist               // Creates an Array with distribution (uint16 dist type index). Two stack operands are prng source (top), shape.
	OpTensor                 // Creates a new Tensor. Two stack operands are shape, array (top).
	OpAdd                    // Adds two stack operands.
)

////////////////////////////////////////////////////////////////////////////////////////////////////
// OpCode

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
	OpLoadG:    {"OpLoadG", []int{2}},
	OpStoreG:   {"OpStoreG", []int{2}},
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
// Instructions

// Prints all instructions in disassembly form.
//
// The basic format is (`%6d %10s %s`, address, opcode, operands).
func (ins Instructions) String() string {
	return ins.DebugString(0 /*startIndex*/, -1 /*numInstructions*/)
}

// Prints all instructions.
//
// With startIndex and numInstructions, it is useful to print a slice of the full program.
// If numInstructions == -1, it means printing all.
func (ins Instructions) DebugString(startIndex int, numInstructions int) string {
	var buf bytes.Buffer

	numPrinted := 0

	i := 0
	for i < len(ins) {
		def, err := Lookup(Opcode(ins[i]))
		if err != nil {
			fmt.Fprintf(&buf, "error raised during printing instructions: %v\naborted...\n", err)
			break
		}

		operands, offset := readOperands(def, ins[i+1:])
		fmt.Fprintf(&buf, "%06d %s\n", i+startIndex, fmtInstruction(def, operands))

		numPrinted++
		if numInstructions != -1 && numPrinted >= numInstructions {
			break
		}

		i += 1 + offset
	}

	return buf.String()
}

// Returns list of operands, and offset for next Instruction in `ins`.
func readOperands(def *Definition, ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths))
	offset := 0

	for i, w := range def.OperandWidths {
		switch w {
		case 2:
			operands[i] = int(ReadUint16(ins[offset:]))
		default:
			panic(fmt.Sprintf("unsupported width (%v) for op: %v", w, def.Name))
		}

		offset += w
	}

	return operands, offset
}

func fmtInstruction(def *Definition, operands []int) string {
	count := len(def.OperandWidths)
	if count != len(operands) {
		panic("internal error: operands count mismatch with opcode definition.")
	}

	switch count {
	case 0:
		return fmt.Sprintf("%-10s", def.Name)
	case 1:
		return fmt.Sprintf("%-10s %d", def.Name, operands[0])
	default:
		panic(fmt.Sprintf("unsupported op count for formatting (%v) for op: %v", count, def.Name))
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Public Helper Methods

func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}
