package code

import (
	"github.com/xiejw/mlvm/go/object"
)

const defaultProgramInitSize = 128

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
