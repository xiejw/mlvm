package compiler

import (
	"strings"
	"testing"

	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/syntax/ast"
)

func assertNoErr(t *testing.T, err *errors.DiagnosisError) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func assertInstructions(t *testing.T, expected, got string) {
	t.Helper()
	expected = strings.Trim(expected, "\n")
	got = strings.Trim(got, "\n")
	if expected != got {
		t.Fatalf(
			"instructions mismatch. "+
				"expected:\n======\n%v\n======\n"+
				"got:\n=====\n%v\n=====",
			expected, got)
	}
}

func makeOpHelper(t *testing.T, op code.Opcode, args ...int) []byte {
	t.Helper()
	ins, err := code.MakeOp(op, args...)
	if err != nil {
		t.Fatalf("failed to make op: %v", err)
	}

	return ins
}

func TestExprSingleLiteral(t *testing.T) {
	c, err := Compile(&ast.Program{
		Expressions: []ast.Expr{
			&ast.IntegerLiteral{123},
		},
	})
	assertNoErr(t, err)
	got := c.Instructions.String()

	var ins code.Instructions
	ins = append(ins, makeOpHelper(t, code.OpConstant, 0)...)
	expected := ins.String()

	assertInstructions(t, expected, got)
}

func TestExprTensorStoreLoad(t *testing.T) {
	c, err := Compile(&ast.Program{
		Expressions: []ast.Expr{
			&ast.FunctionCall{
				Func: &ast.Id{"store_load"},
				Args: []ast.Expr{
					&ast.StringLiteral{"a"},
				},
			},
		},
	})
	assertNoErr(t, err)
	got := c.Instructions.String()

	var ins code.Instructions
	ins = append(ins, makeOpHelper(t, code.OpConstant, 0)...)
	ins = append(ins, makeOpHelper(t, code.OpLoadT)...)
	expected := ins.String()

	assertInstructions(t, expected, got)
}
