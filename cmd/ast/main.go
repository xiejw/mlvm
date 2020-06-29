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
			Name: &ast.Identifier{"prng_dist"},
			Args: []ast.Expression{
				&ast.FunctionCall{
					Name: &ast.Identifier{"prng_new"},
					Args: []ast.Expression{
						&ast.IntegerLiteral{123}, // seed
					},
				},
				&ast.IntegerLiteral{1}, // dist type
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
	err = m.Run()
	if err != nil {
		log.Fatalf("failed to run vm: %v", err)
	}

	log.Printf("Results:\n%v\n", m.StackTop())
}
