package compiler

import (
	"fmt"

	"github.com/xiejw/mlvm/go/syntax/ast"
)

func checkSingleArg(fname string, args []ast.Expression) (ast.Expression, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("function (\"%v\") should have exactly 1 arg.", fname)
	}
	return args[0], nil
}

func (b *Builder) compileBuiltinFn(fn *ast.FunctionCall) error {
	fnName := fn.Name.Value

	switch fnName {
	case "store_load":
		arg, err := checkSingleArg(fnName, fn.Args)
		if err != nil {
			return err
		}

		err = b.compileExpression(arg)
		if err != nil {
			return err
		}
		b.emitLoadTensor()
		return nil
	default:
		return fmt.Errorf("unsupported function name: %v", fnName)
	}
}
