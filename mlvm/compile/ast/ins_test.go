package ast

import (
	"testing"

	"github.com/xiejw/mlvm/mlvm/array"
)

func TestInstructionAdd(t *testing.T) {
	a := array.NewArrayOrDie("a", []array.Dimension{2, 1}, []array.Float{1.0, 2.0})
	ta := newConstantTensor(a)

	ins := newInstruction("add", OpAdd(), ta, ta)

	got := ins.String()
	expected := `Ins{"add", (Constant{"a", <2, 1>}, Constant{"a", <2, 1>}) -> (Result{"%o_0", <2, 1>})}`
	if expected != got {
		t.Errorf("Ins mismatch: Expected: %v, Got: %v", expected, got)
	}
}
