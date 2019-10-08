package ast

import (
	"testing"
)

func TestOpAdd(t *testing.T) {
	op := OpAdd()
	if op.Kind() != OpKAdd {
		t.Errorf("Kind mismatch.")
	}
}
