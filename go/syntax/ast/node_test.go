package ast

import (
	"strings"
	"testing"
)

func assertAstOutput(t *testing.T, ast *Program, expected string) {
	t.Helper()
	expected = strings.Trim(expected, "\n")
	got := strings.Trim(ast.Expressions.String(), "\n")
	if expected != got {
		t.Errorf("ast mismatch. expected: `%v`, got: `%v`.", expected, got)
	}
}

func makeSingleExprProgram(expr Expression) *Program {
	return &Program{Expressions: []Expression{expr}}
}

func TestIntegerLiteral(t *testing.T) {
	p := makeSingleExprProgram(&IntegerLiteral{123})
	assertAstOutput(t, p, `Int(123)`)
}

func TestFloatLiteral(t *testing.T) {
	p := makeSingleExprProgram(&FloatLiteral{98.76})
	assertAstOutput(t, p, `Float(98.76)`)
}

func TestStringLiteral(t *testing.T) {
	p := makeSingleExprProgram(&StringLiteral{"abc"})
	assertAstOutput(t, p, `Str("abc")`)
}

func TestFuncCall(t *testing.T) {
	p := makeSingleExprProgram(&FunctionCall{
		Func: &Identifier{"fn_name"},
		Args: []Expression{&Identifier{"arg_0"}},
	})
	assertAstOutput(t, p, `Func(ID(fn_name), ID(arg_0))`)
}
