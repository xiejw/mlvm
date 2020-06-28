package main

import (
	"log"

	"github.com/xiejw/mlvm/go/compiler"
	"github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/syntax/ast"
	"github.com/xiejw/mlvm/go/vm"
)

func main() {

	statements := make([]ast.Statement, 0)
	statements = append(statements, &ast.ExprStatement{
		Value: &ast.FunctionCall{
			Name: &ast.Identifier{"store_load"},
			Args: []ast.Expression{
				&ast.StringLiteral{"a"},
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

	ts := vm.NewTensorStore()
	shape := object.NewShape([]object.NamedDim{{"x", 2}})
	array := &object.Array{[]float32{1.0, 2.0}}
	tensor := &object.Tensor{shape, array}
	err = ts.Store("a", tensor)
	if err != nil {
		log.Fatalf("failed to store tensor into tensor store: %v", err)
	}

	m := vm.NewVMWithTensorStore(o, ts)

	log.Printf("Running VM\n")
	err = m.Run()
	if err != nil {
		log.Fatalf("failed to run vm: %v", err)
	}

	log.Printf("Results:\n%v\n", m.StackTop())
}
