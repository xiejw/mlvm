package main

import (
	"log"

	"github.com/xiejw/mlvm/go/compiler"
	"github.com/xiejw/mlvm/go/syntax/ast"
	"github.com/xiejw/mlvm/go/vm"
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

	log.Printf("Compiled Code:\n%v\n", o.Instructions)

	m := vm.NewVM(o)

	log.Printf("Running VM\n")
	m.Run()

	log.Printf("Results:\n%v\n", m.StackTop())
}
