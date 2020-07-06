package parser

import (
	"strings"
	"testing"

	"github.com/xiejw/mlvm/go/syntax/ast"
)

func assertNoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func assertAstOutput(t *testing.T, ast *ast.Program, expected string) {
	t.Helper()
	expected = strings.Trim(expected, "\n")
	got := strings.Trim(ast.Expressions.String(), "\n")
	if expected != got {
		t.Errorf("ast mismatch. expected: `%v`, got: `%v`.", expected, got)
	}
}

func TestIntegerLiteral(t *testing.T) {
	p := New([]byte("123"))
	ast, err := p.ParseAst()
	assertNoErr(t, err)
	assertAstOutput(t, ast, "Int(123)")
}

func TestFunctionCall(t *testing.T) {
	p := New([]byte("(+ a 123)"))
	ast, err := p.ParseAst()
	assertNoErr(t, err)
	assertAstOutput(t, ast, "Func(ID(+), ID(a), Int(123))")
}

func TestNestedFunctionCall(t *testing.T) {
	p := New([]byte("(+ a (call b a))"))
	ast, err := p.ParseAst()
	assertNoErr(t, err)
	assertAstOutput(t, ast, `Func(ID(+), ID(a), Func(ID(call), ID(b), ID(a)))`)
}
