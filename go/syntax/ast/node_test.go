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

func TestIntLiteral(t *testing.T) {
	p := &Program{
		Expressions: []Expression{
			&IntegerLiteral{123},
		},
	}

	if len(p.Expressions) != 1 {
		t.Errorf("expect single expr")
	}

	assertAstOutput(t, p, `Int(123)`)
}

func TestStrLiteral(t *testing.T) {
	p := &Program{
		Expressions: []Expression{
			&StringLiteral{"abc"},
		},
	}

	if len(p.Expressions) != 1 {
		t.Errorf("expect single expr")
	}

	assertAstOutput(t, p, `Str("abc")`)
}

func TestFuncCall(t *testing.T) {
	p := &Program{
		Expressions: []Expression{
			&FunctionCall{
				Func: &Identifier{"fn_name"},
				Args: []Expression{&Identifier{"arg_0"}},
			},
		},
	}

	if len(p.Expressions) != 1 {
		t.Errorf("expect single expr")
	}

	assertAstOutput(t, p, `Func(ID(fn_name), ID(arg_0))`)
}
