package code

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Instructions []byte

type Opcode byte

const (
	OpConstant Opcode = iota
	OpTensor
)

type Definition struct {
	Name          string
	OperandWidths []int
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// OpCode

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
	OpTensor:   {"OpTensor", []int{}},
}

func Lookup(op Opcode) (*Definition, error) {
	def, ok := definitions[op]
	if !ok {
		return nil, fmt.Errorf("Opcode %d undefined", op)
	}

	return def, nil
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
		return nil, fmt.Errorf("Operand counts mismatch: expected %v, got %v",
			len(def.OperandWidths), len(operands))
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

func (ins Instructions) String() string {
	var buf bytes.Buffer

	i := 0
	for i < len(ins) {
		def, err := Lookup(Opcode(ins[i]))
		if err != nil {
			fmt.Fprintf(&buf, "error raised during printing instructions: %v\naborted...\n", err)
			break
		}

		operands, offset := readOperands(def, ins[i+1:])
		fmt.Fprintf(&buf, "%06d %s\n", i, fmtInstruction(def, operands))
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
		return fmt.Sprintf("%s", def.Name)
	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0])
	default:
		panic(fmt.Sprintf("unsupported op count for formatting (%v) for op: %v", count, def.Name))
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Public Helper Methods

func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}
