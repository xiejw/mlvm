package code

const defaultProgramInitSize = 128

type Program struct {
	Instructions Instructions
	Constants    []interface{} // change to object.
}

func NewProgram() *Program {
	return &Program{
		Instructions: make([]byte, 0, defaultProgramInitSize),
		Constants:    make([]interface{}, 0),
	}
}
