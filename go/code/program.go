package code

import (
	"bytes"
	"fmt"
	"io"

	"github.com/xiejw/mlvm/go/object"
)

const defaultProgramInitSize = 128

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
