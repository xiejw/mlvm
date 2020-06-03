package ast

import (
	"bytes"
	"testing"

	"github.com/xiejw/mlvm/go/object"
)

func TestDeclNamedDim(t *testing.T) {
	// @batch = 32;
	decl := &Decl{
		Name: "@batch",
		Value: &Expression{
			Type:  ExTpValue,
			Value: &object.Integer{32},
		},
	}

	var buf bytes.Buffer
	decl.WriteDebugString(&buf, "")
	outs := buf.String()

	expected := `Decl{
  Token: "@batch"
  Value: Expr{
    Value: Int(32)
  }
}
`
	if outs != expected {
		t.Errorf("Mismatches. expected:\n%v\ngot:\n%v\n", expected, outs)
	}
}
