package code

import (
	"fmt"

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
	return fmt.Sprintf("-> Constants:\n\n%v\n\n-> Instruction:\n%v\n", p.Constants, p.Instructions)
}
