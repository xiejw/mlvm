package compiler

import (
	"fmt"

	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/syntax/ast"
)

// Compiles ast.Program to code.Program.
func Compile(src *ast.Program) (*code.Program, error) {

	b := NewBuilder(src)

	for _, src_statement := range src.Statements {
		err := b.compileStatement(src_statement)
		if err != nil {
			return nil, err
		}
	}

	return b.CompiledCode(), nil
}

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
	return nil
}

func (b *Builder) CompiledCode() *code.Program {
	return b.output
}

func (b *Builder) emitIntConstant(literal *ast.IntegerLiteral) int {
	return 0
}

func (b *Builder) emitLoadConstant(constIndex int) {
}

func (b *Builder) compileStatement(statement ast.Statement) error {
	switch v := statement.(type) {
	case *ast.ExprStatement:
		err := b.compileExpression(v.Value)
		return fmt.Errorf("error during compiling expr statement: %w", err)
	default:
		return fmt.Errorf("unsupported statement.")
	}
}

func (b *Builder) compileExpression(expr ast.Expression) error {
	switch v := expr.(type) {
	case *ast.IntegerLiteral:
		index := b.emitIntConstant(v)
		b.emitLoadConstant(index)
		return nil
	default:
		return fmt.Errorf("unsupported statement.")
	}
	return nil
}
