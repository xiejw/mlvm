package tests

import (
	"reflect"
	"testing"

	"github.com/xiejw/mlvm/go/compiler"
	"github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/syntax/ast"
	"github.com/xiejw/mlvm/go/vm"
)

func TestSingleExpression(t *testing.T) {
	statements := []ast.Statement{
		&ast.ExprStatement{
			Value: &ast.IntegerLiteral{123},
		}}

	p := &ast.Program{
		Statements: statements,
	}

	o, err := compiler.Compile(p)
	if err != nil {
		t.Fatalf("unexpected compile error: %v", err)
	}

	m := vm.NewVM(o)
	m.Run()

	expected := &object.Integer{123}
	got := m.StackTop()

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("unexpected output.")
	}
}
