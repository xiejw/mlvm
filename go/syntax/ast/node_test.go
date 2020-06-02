package ast

import (
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

	if decl.Name != "@batch" {
		t.Errorf("name mismatches.")
	}
}
