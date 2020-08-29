// Package `code` provides the compiled program to run with vm.
package code

import (
	"bytes"
	"fmt"
	"io"

	"github.com/xiejw/mlvm/go/object"
)

const defaultProgramInitSize = 128

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

func (cs Constants) String() string {
	var buf bytes.Buffer
	cs.ToHumanReadableString(&buf)
	return buf.String()
}

func (cs Constants) ToHumanReadableString(w io.Writer) {
	if len(cs) == 0 {
		fmt.Fprintf(w, "(empty)")
		return
	}

	fmt.Fprint(w, "[\n")
	for i, c := range cs {
		fmt.Fprintf(w, "  %3d: %v\n", i, c)
	}
	fmt.Fprint(w, "]")
}

func (p *Program) String() string {
	var buf bytes.Buffer
	p.ToHumanReadableString(&buf)
	return buf.String()
}

func (p *Program) ToHumanReadableString(w io.Writer) {
	fmt.Fprint(w, "-> Constants:\n\n")
	p.Constants.ToHumanReadableString(w)
	fmt.Fprintf(w, "\n\n-> Instruction:\n\n%v\n",
		p.Instructions)
}
