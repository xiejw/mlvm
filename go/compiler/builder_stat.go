package compiler

import (
	"fmt"

	"github.com/xiejw/mlvm/go/syntax/ast"
)

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
