package compiler

import (
	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/syntax/ast"
)

// Compiles the expression.
func (b *Builder) compileExpression(expr ast.Expression) *errors.DiagnosisError {
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
		err := b.compileBuiltinFn(v)
		if err != nil {
			return err.EmitDiagnosisNote(
				"compiling function call. currently only support " +
					"built-in functions")
		}
		return nil

	default:
		return errors.NewDiagnosisError(
			"unsupported expression to be compiled. currently "+
				"only support integer literal, string literal, "+
				"function call. got: %+v",
			ast.Expressions([]ast.Expression{expr}))
	}
}
