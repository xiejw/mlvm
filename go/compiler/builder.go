package compiler

import (
	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/syntax/ast"
)

type Builder struct {
	// Interal fields
	input  *ast.Program
	output *code.Program
}

func NewBuilder(src *ast.Program) *Builder {
	return &Builder{
		input: src,
		output: &code.Program{
			Instructions: make([]byte, 0),
			Constants:    make([]object.Object, 0),
		},
	}
}

func (b *Builder) Compile() error {

	// statements := b.input.Statements
	// finalStatementIndex := len(statements) - 1

	// for i, src_statement := range statements {
	// 	err := b.compileStatement(src_statement)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if _, ok := src_statement.(*ast.ExprStatement); ok && i != finalStatementIndex {
	// 		b.emitPop()
	// 	}
	// }
	return nil
}

func (b *Builder) CompiledCode() *code.Program {
	return b.output
}
