package ast

import (
	"strings"
	"testing"
)

func assertAstOutput(t *testing.T, ast *Program, expected string) {
	t.Helper()
	expected = strings.Trim(expected, "\n")
	got := strings.Trim(ast.Exprs.String(), "\n")
	if expected != got {
		t.Errorf("ast mismatch. expected: `%v`, got: `%v`.", expected, got)
	}
}

func makeSingleExprProgram(expr Expr) *Program {
	return &Program{Exprs: []Expr{expr}}
}

func TestIdentifier(t *testing.T) {
	p := makeSingleExprProgram(&Id{Value: "abc"})
	assertAstOutput(t, p, `Id(abc)`)
}

func TestIntegerLit(t *testing.T) {
	p := makeSingleExprProgram(&IntLit{Value: 123})
	assertAstOutput(t, p, `Int(123)`)
}

func TestFloatLit(t *testing.T) {
	p := makeSingleExprProgram(&FloatLit{Value: 98.76})
	assertAstOutput(t, p, `Float(98.76)`)
}

func TestShapeLit(t *testing.T) {
	p := makeSingleExprProgram(&ShapeLit{
		Dims: []*Id{
			&Id{Value: "@a"},
			&Id{Value: "@b"},
		},
	})
	assertAstOutput(t, p, `Shape(Id(@a), Id(@b))`)
}

func TestArrayLit(t *testing.T) {
	p := makeSingleExprProgram(&ArrayLit{
		Values: []*FloatLit{
			&FloatLit{Value: 1.76},
			&FloatLit{Value: 2.98},
		},
	})
	assertAstOutput(t, p, `Array(Float(1.76), Float(2.98))`)
}

func TestStringLit(t *testing.T) {
	p := makeSingleExprProgram(&StringLit{Value: "abc"})
	assertAstOutput(t, p, `String("abc")`)
}

func TestFuncCall(t *testing.T) {
	p := makeSingleExprProgram(&App{
		Func: &Id{Value: "fn_name"},
		Args: []Expr{&Id{Value: "arg_0"}},
	})
	assertAstOutput(t, p, `App(Id(fn_name), Id(arg_0))`)
}
