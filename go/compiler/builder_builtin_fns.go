package compiler

import (
	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/syntax/ast"
)

func checkSingleArg(fname string, args []ast.Expression) (
	ast.Expression, *errors.DiagnosisError,
) {
	if len(args) != 1 {
		return nil, errors.NewDiagnosisError(
			"function (\"%v\") should have exactly 1 arg, got: %v.",
			fname, len(args))
	}
	return args[0], nil
}

func (b *Builder) compileBuiltinFn(fn *ast.FunctionCall) *errors.DiagnosisError {
	fnName := fn.Func.Value

	switch fnName {
	case "store_load":
		arg, err := checkSingleArg(fnName, fn.Args)
		if err != nil {
			return err.EmitDiagnosisNote("compiling built-in function \"%v\"", fnName)
		}

		err = b.compileExpression(arg)
		if err != nil {
			return err.EmitDiagnosisNote(
				"compiling arguments for built-in function \"%v\"", fnName)
		}
		b.emitLoadTensor()
		return nil
	default:
		return errors.NewDiagnosisError("unsupported built-in function name: %v", fnName)
	}
}
