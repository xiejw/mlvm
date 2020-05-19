package code

import (
	"testing"
)

func checkNoErr(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("Did not expect error.")
	}
}

func TestOpcodes(t *testing.T) {
	ops := []struct {
		op   Opcode
		name string
	}{
		{OpConstant, "OpConstant"},
	}

	for _, testOp := range ops {
		op, err := Lookup(testOp.op)
		checkNoErr(t, err)
		if op.Name != testOp.name {
			t.Fatalf("Op name mistmatch")
		}
	}
}
