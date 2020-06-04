package ast

import (
	"bytes"
	"testing"

	"github.com/xiejw/mlvm/go/object"
)

func assertStringRepr(t *testing.T, expectedStartedWithNewLine, got string) {

	expected := expectedStartedWithNewLine[1:]
	if got[:] != expected[:] {
		t.Errorf("String representation mismatches.\n-> expected:\n%v\n-> got:\n%v\n",
			expected, got)
	}
}

func TestDeclNamedDim(t *testing.T) {
	// let @batch = 32;
	decl := &Decl{
		Name: "@batch",
		Type: &Type{Kind: TpKdNamedDim},
		Value: &Expression{
			Type:  ExTpValue,
			Value: &object.Integer{32},
		},
	}

	var buf bytes.Buffer
	decl.WriteDebugString(&buf, "")
	got := buf.String()

	expected := `
Decl{
  Name: "@batch"
  Type: Type(NamedDim)
  Value: Expr{
    Value: Int(32)
  }
}
`

	assertStringRepr(t, expected, got)
}
