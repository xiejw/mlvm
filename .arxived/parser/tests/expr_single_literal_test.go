package tests

import (
	"reflect"
	"testing"

	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/compiler"
	"github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/syntax/parser"
	"github.com/xiejw/mlvm/go/vm"
)

func assertNoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func assertNoDiagnosisError(t *testing.T, err *errors.DError) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func assertSingleOutput(t *testing.T, outputs vm.Outputs, err error) object.Object {
	t.Helper()
	assertNoErr(t, err)
	if len(outputs) != 1 {
		t.Fatalf("unexpected single output, got: %v", outputs)
	}
	return outputs[0]
}

func TestExprSingleLiteral(t *testing.T) {
	p, diagnosisError := parser.New([]byte("123")).ParseAst()
	assertNoDiagnosisError(t, diagnosisError)

	o, diagnosisError := compiler.Compile(p)
	assertNoDiagnosisError(t, diagnosisError)

	m := vm.NewVM(o)
	outputs, err := m.Run()
	got := assertSingleOutput(t, outputs, err)

	expected := &object.Integer{123}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("unexpected output.")
	}
}
