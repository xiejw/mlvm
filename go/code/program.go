// Package `code` provides the compiled program to run with vm.
package code

import (
	"bytes"
	"fmt"
	"io"

	"github.com/xiejw/mlvm/go/object"
)

const defaultProgramInitSize = 128

type Instructions []byte
type Constants []object.Object

// All fields are not intended to be mutated. So the program can be re-used.
type Program struct {
	Instructions Instructions
	Constants    Constants
}

func NewProgram() *Program {
	return &Program{
		Instructions: make([]byte, 0, defaultProgramInitSize),
		Constants:    make([]object.Object, 0),
	}
}

///////////////////////////////////////////////////////////////////////////////
// String related.
///////////////////////////////////////////////////////////////////////////////

func (p *Program) String() string {
	var buf bytes.Buffer
	p.DebugString(&buf)
	return buf.String()
}

func (p *Program) DebugString(w io.Writer) {
	fmt.Fprint(w, "-> Constants:\n\n")
	p.Constants.DebugString(w)
	fmt.Fprintf(w, "\n-> Instruction:\n\n%v\n",
		p.Instructions)
}

func (cs Constants) String() string {
	var buf bytes.Buffer
	cs.DebugString(&buf)
	return buf.String()
}

func (cs Constants) DebugString(w io.Writer) {
	if len(cs) == 0 {
		fmt.Fprintf(w, "(empty)\n")
		return
	}

	fmt.Fprint(w, "[\n")
	for i, c := range cs {
		fmt.Fprintf(w, "  %3d: %v\n", i, c)
	}
	fmt.Fprint(w, "]\n")
}

// Prints all instructions in disassembly form.
//
// The basic format is (`%6d %10s %s`, address, opcode, operands).
func (ins Instructions) String() string {
	return ins.DebugString(0 /*startIndex*/, -1 /*numInstructions*/)
}

// Prints all instructions.
//
// With startIndex and numInstructions, it is useful to print a slice of the full program.  If
// numInstructions == -1, it means printing all.
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
		return fmt.Sprintf("%s", def.Name)
	case 1:
		return fmt.Sprintf("%-10s %d", def.Name, operands[0])
	default:
		panic(fmt.Sprintf("unsupported op count for formatting (%v) for op: %v", count, def.Name))
	}
}
