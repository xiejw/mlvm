package compiler

import (
	"strings"
	"testing"

	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/syntax/ast"
)

func assertNoErr(t *testing.T, err error) {
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
	assertNoErr(t, err)

	return ins
}

func TestExprSingleLiteral(t *testing.T) {
	c, err := Compile(&ast.Program{
		Expressions: []ast.Expression{
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
		Expressions: []ast.Expression{
			&ast.FunctionCall{
				Func: &ast.Identifier{"store_load"},
				Args: []ast.Expression{
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
