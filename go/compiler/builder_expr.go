package compiler

import (
	"fmt"

	"github.com/xiejw/mlvm/go/syntax/ast"
)

// Compiles the expression.
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
		return fmt.Errorf("compiler error: unsupported expression: %+v", ast.Expressions([]ast.Expression{expr}))
	}
}
