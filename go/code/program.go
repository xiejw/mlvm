package code

const defaultProgramInitSize = 128

type Program struct {
	Instructions Instructions
	Constants    []Object
}

func NewProgram() *Program {
	return &Program{
		Instructions: make([]byte, 0, defaultProgramInitSize),
		Constants:    make([]Object, 0),
	}
}
