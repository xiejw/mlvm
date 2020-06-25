package main

import (
	"fmt"
	"log"
	// "encoding/json"
	"github.com/xiejw/mlvm/go/compiler"
	"github.com/xiejw/mlvm/go/syntax/ast"
)

func main() {

	statements := make([]ast.Statement, 0)
	// statements = append(statements, &ast.DeclStatement{
	// 	Name: &ast.Identifier{"seed"},
	// 	Value: &ast.FunctionCall{
	// 		Name: &ast.Identifier{"prng_new"},
	// 		Args: []ast.Expression{
	// 			&ast.IntegerLiteral{123},
	// 		},
	// 	},
	// })
	statements = append(statements, &ast.ExprStatement{
		Value: &ast.IntegerLiteral{123},
	})

	p := &ast.Program{
		Statements: statements,
	}

	o, err := compiler.Compile(p)
	if err != nil {
		log.Fatalf("Failed to compile %v", err)
	}

	fmt.Printf("%v\n", o.Instructions)
}
