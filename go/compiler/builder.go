package compiler

import (
	"fmt"

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

	statements := b.input.Statements
	finalStatementIndex := len(statements) - 1

	for i, src_statement := range statements {
		err := b.compileStatement(src_statement)
		if err != nil {
			return err
		}
		if _, ok := src_statement.(*ast.ExprStatement); ok && i != finalStatementIndex {
			b.emitPop()
		}
	}
	return nil
}

func (b *Builder) CompiledCode() *code.Program {
	return b.output
}

func (b *Builder) compileStatement(statement ast.Statement) error {
	switch v := statement.(type) {
	case *ast.ExprStatement:
		err := b.compileExpression(v.Value)
		if err != nil {
			return fmt.Errorf("error during compiling expr statement: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("unsupported statement.")
	}
}

func (b *Builder) compileExpression(expr ast.Expression) error {
	switch v := expr.(type) {
	case *ast.IntegerLiteral:
		index := b.emitIntegerConstant(v)
		b.emitLoadConstant(index)
		return nil
	case *ast.StringLiteral:
		index := b.emitStringConstant(v)
		b.emitLoadConstant(index)
		return nil
	case *ast.FunctionCall:
		// Currently only supports limited bultin-ins.
		return b.compileBuiltinFn(v)

	default:
		return fmt.Errorf("unsupported expression: %+v", expr)
	}
	return nil
}
