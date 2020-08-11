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

func makeSingleExprProgram(expr Expr) *Program {
	return &Program{Expressions: []Expr{expr}}
}

func TestIdentifier(t *testing.T) {
	p := makeSingleExprProgram(&Id{"abc"})
	assertAstOutput(t, p, `ID(abc)`)
}

func TestIntegerLit(t *testing.T) {
	p := makeSingleExprProgram(&IntLit{123})
	assertAstOutput(t, p, `Int(123)`)
}

func TestFloatLit(t *testing.T) {
	p := makeSingleExprProgram(&FloatLit{98.76})
	assertAstOutput(t, p, `Float(98.76)`)
}

func TestShapeLit(t *testing.T) {
	p := makeSingleExprProgram(&ShapeLit{
		[]*Id{
			&Id{"@a"},
			&Id{"@b"},
		},
	})
	assertAstOutput(t, p, `Shape(ID(@a), ID(@b))`)
}

func TestArrayLit(t *testing.T) {
	p := makeSingleExprProgram(&ArrayLit{
		[]*FloatLit{
			&FloatLit{1.76},
			&FloatLit{2.98},
		},
	})
	assertAstOutput(t, p, `Array(Float(1.76), Float(2.98))`)
}

func TestStringLit(t *testing.T) {
	p := makeSingleExprProgram(&StringLit{"abc"})
	assertAstOutput(t, p, `Str("abc")`)
}

func TestFuncCall(t *testing.T) {
	p := makeSingleExprProgram(&FunctionCall{
		Func: &Id{"fn_name"},
		Args: []Expr{&Id{"arg_0"}},
	})
	assertAstOutput(t, p, `Func(ID(fn_name), ID(arg_0))`)
}
