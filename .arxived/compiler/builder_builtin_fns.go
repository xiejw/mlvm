package compiler

import (
	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/syntax/ast"
)

func checkSingleArg(fname string, args []ast.Expr) (
	ast.Expr, *errors.DError,
) {
	if len(args) != 1 {
		return nil, errors.New(
			"function (\"%v\") should have exactly 1 arg, got: %v.",
			fname, len(args))
	}
	return args[0], nil
}

func checkDoubleArgs(fname string, args []ast.Expr) (
	ast.Expr, ast.Expr, *errors.DError,
) {
	if len(args) != 2 {
		return nil, nil, errors.New(
			"function (\"%v\") should have exactly 2 args, got: %v.",
			fname, len(args))
	}
	return args[0], args[1], nil
}

func (b *Builder) compileBuiltinFn(fn *ast.App) *errors.DError {
	fnName := fn.Func.Value

	switch fnName {
	case "store_load":
		arg, err := checkSingleArg(fnName, fn.Args)
		if err != nil {
			return err.EmitNote("compiling built-in function \"%v\"", fnName)
		}

		err = b.compileExpression(arg)
		if err != nil {
			return err.EmitNote(
				"compiling the argument for built-in function \"%v\"", fnName)
		}
		b.emitLoadTensor()
		return nil

	case "+":
		arg1, arg2, err := checkDoubleArgs(fnName, fn.Args)
		if err != nil {
			return err.EmitNote("compiling built-in function \"%v\"", fnName)
		}

		err = b.compileExpression(arg2)
		if err != nil {
			return err.EmitNote(
				"compiling the second argument for built-in function \"%v\"", fnName)
		}

		err = b.compileExpression(arg1)
		if err != nil {
			return err.EmitNote(
				"compiling the first argument for built-in function \"%v\"", fnName)
		}
		b.emitAdd()
		return nil
	default:
		return errors.New("unsupported built-in function name: %v", fnName)
	}
}
