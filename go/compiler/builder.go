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

	expressions := b.input.Expressions
	finalStatementIndex := len(expressions) - 1

	for i, expr := range expressions {
		err := b.compileExpression(expr)
		if err != nil {
			return err
		}
		if i != finalStatementIndex {
			b.emitPop()
		}
	}
	return nil
}

func (b *Builder) CompiledCode() *code.Program {
	return b.output
}
