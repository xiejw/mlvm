package tests

import (
	"reflect"
	"testing"

	"github.com/xiejw/mlvm/go/compiler"
	"github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/syntax/ast"
	"github.com/xiejw/mlvm/go/vm"
)

func createSimpleTensor() *object.Tensor {
	shape := object.NewShape([]object.NamedDim{{"x", 2}})
	array := &object.Array{[]float32{1.0, 2.0}}
	return &object.Tensor{shape, array}
}

func TestExprTensorStoreLoad(t *testing.T) {
	statements := []ast.Statement{
		&ast.ExprStatement{
			Value: &ast.FunctionCall{
				Name: &ast.Identifier{"store_load"},
				Args: []ast.Expression{
					&ast.StringLiteral{"a"},
				},
			},
		},
	}

	p := &ast.Program{
		Statements: statements,
	}

	o, err := compiler.Compile(p)
	assertNoErr(t, err)

	ts := vm.NewTensorStore()
	err = ts.Store("a", createSimpleTensor())
	assertNoErr(t, err)

	m := vm.NewVMWithTensorStore(o, ts)

	err = m.Run()
	assertNoErr(t, err)

	expected := createSimpleTensor()
	got := m.StackTop()

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("unexpected output.")
	}
}
