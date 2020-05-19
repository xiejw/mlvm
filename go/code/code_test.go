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

func TestInstructionString(t *testing.T) {
	expected := "000000 OpConstant 123\n"

	ins, err := MakeOp(OpConstant, 123)
	checkNoErr(t, err)

	got := Instructions(ins).String()

	if expected != got {
		t.Errorf("Unexpected Instructions String(): expected:\n%v\ngot:\n%v\n", expected, got)
	}
}
