package compiler

import (
	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/syntax/ast"
)

// Compiles ast.Program to code.Program.
func Compile(src *ast.Program) (*code.Program, error) {

	b := NewBuilder(src)
	err := b.Compile()
	if err != nil {
		return nil, err
	}

	return b.CompiledCode(), nil
}
