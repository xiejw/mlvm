package code

import (
	"github.com/xiejw/mlvm/go/object"
)

const defaultProgramInitSize = 128

// All fields are not intended to be mutated. So the program can be re-used.
type Program struct {
	Instructions Instructions
	Constants    []object.Object
}

func NewProgram() *Program {
	return &Program{
		Instructions: make([]byte, 0, defaultProgramInitSize),
		Constants:    make([]object.Object, 0),
	}
}
