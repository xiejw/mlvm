package compiler

import (
	"fmt"

	_ "github.com/xiejw/mlvm/go/code"
	_ "github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/syntax/ast"
)

func (b *Builder) compileBuiltinFn(fn *ast.FunctionCall) error {
	fnName := fn.Name.Value

	switch fnName {
	case "store_load":
		if len(fn.Args) != 1 {
			return fmt.Errorf("function (\"%v\") should have exactly 1 arg.", fnName)
		}
		err := b.compileExpression(fn.Args[0])
		if err != nil {
			return err
		}
		fmt.Printf("TODO support Load store.\n")
		return nil
	default:
		return fmt.Errorf("unsupported function name: %v", fnName)
	}
}
