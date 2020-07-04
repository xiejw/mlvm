package main

import (
	"log"

	_ "github.com/xiejw/mlvm/go/compiler"
	_ "github.com/xiejw/mlvm/go/syntax/ast"
	_ "github.com/xiejw/mlvm/go/vm"
)

func main() {
	log.Printf("Hello")

	// 	statements := make([]ast.Statement, 0)
	// 	statements = append(statements, &ast.ExprStatement{
	// 		Value: &ast.FunctionCall{
	// 			Name: &ast.Identifier{"prng_dist"},
	// 			Args: []ast.Expression{
	// 				&ast.FunctionCall{
	// 					Name: &ast.Identifier{"prng_new"},
	// 					Args: []ast.Expression{
	// 						&ast.IntegerLiteral{123}, // seed
	// 					},
	// 				},
	// 				&ast.IntegerLiteral{1}, // dist type
	// 			},
	// 		},
	// 	})
	//
	// 	p := &ast.Program{
	// 		Statements: statements,
	// 	}
	//
	// 	o, err := compiler.Compile(p)
	// 	if err != nil {
	// 		log.Fatalf("failed to compile: %v", err)
	// 	}
	//
	// 	log.Printf("Compiled Code:\n\n%v", o)
	//
	// 	m := vm.NewVM(o)
	//
	// 	log.Printf("Running VM\n")
	// 	outputs, err := m.Run()
	// 	if err != nil {
	// 		log.Fatalf("failed to run vm: %v", err)
	// 	}
	//
	// 	log.Printf("Results:\n%v\n", outputs)
}
