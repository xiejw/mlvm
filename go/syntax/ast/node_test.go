package ast

import "testing"

func TestSingleExprStatement(t *testing.T) {
	p := Program{
		Statements: []Statement{
			&ExprStatement{
				Value: &IntLiteral{123},
			},
		},
	}

	if len(p.Statements) != 1 {
		t.Errorf("expect empty Program")
	}
}
