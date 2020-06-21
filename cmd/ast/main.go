package main

import (
	"fmt"
	"log"

	"encoding/json"
	"github.com/xiejw/mlvm/go/syntax/ast"
)

func main() {

	statements := make([]ast.Statement, 0)
	statements = append(statements, &ast.DeclStatement{
		Name: &ast.Identifier{"seed"},
		Value: &ast.FunctionCall{
			Name: &ast.Identifier{"prng_new"},
			Args: []ast.Expression{
				&ast.IntLiteral{123},
			},
		},
	})

	p := &ast.Program{
		Statements: statements,
	}

	prettyJSON, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	fmt.Printf("%s\n", prettyJSON)
}
