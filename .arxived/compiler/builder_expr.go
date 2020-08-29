package compiler

import (
	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/syntax/ast"
)

// Compiles the expression.
func (b *Builder) compileExpression(expr ast.Expr) *errors.DError {
	switch v := expr.(type) {
	case *ast.IntLit:
		index := b.emitIntegerConstant(v)
		b.emitLoadConstant(index)
		return nil
	case *ast.StringLit:
		index := b.emitStringConstant(v)
		b.emitLoadConstant(index)
		return nil
	case *ast.App:
		// Currently only supports limited bultin-ins.
		err := b.compileBuiltinFn(v)
		if err != nil {
			return err.EmitNote(
				"compiling function call. currently only support " +
					"built-in functions")
		}
		return nil

	default:
		return errors.New(
			"unsupported expression to be compiled. currently "+
				"only support integer literal, string literal, "+
				"function call. got: %+v",
			ast.Exprs([]ast.Expr{expr}))
	}
}
