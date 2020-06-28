package main

import (
	"log"

	"github.com/xiejw/mlvm/go/compiler"
	"github.com/xiejw/mlvm/go/syntax/ast"
	"github.com/xiejw/mlvm/go/vm"
)

func main() {

	statements := make([]ast.Statement, 0)
	statements = append(statements, &ast.ExprStatement{
		Value: &ast.FunctionCall{
			Name: &ast.Identifier{"store_load"},
			Args: []ast.Expression{
				&ast.IntegerLiteral{123},
			},
		},
	})

	p := &ast.Program{
		Statements: statements,
	}

	o, err := compiler.Compile(p)
	if err != nil {
		log.Fatalf("failed to compile: %v", err)
	}

	log.Printf("Compiled Code:\n\n%v", o)

	m := vm.NewVM(o)

	log.Printf("Running VM\n")
	m.Run()

	log.Printf("Results:\n%v\n", m.StackTop())
}
